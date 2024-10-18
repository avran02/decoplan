package hub

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"sync"

	"github.com/avran02/decoplan/chat/enum"
	"github.com/avran02/decoplan/chat/internal/dto"
	"github.com/avran02/decoplan/chat/internal/mapper"
	"github.com/avran02/decoplan/chat/internal/models"
	"github.com/avran02/decoplan/chat/internal/service"
	"github.com/gorilla/websocket"

	jsoniter "github.com/json-iterator/go"
)

var (
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type WebsocketHub interface {
	CloseWebsocket(w http.ResponseWriter, r *http.Request)
	RegisterWebsocket(w http.ResponseWriter, r *http.Request)
}

type websocketHub struct {
	clients  map[string]models.WebsocketClient // map[remoteAddr]models.WebsocketClient
	clients2 map[string]string                 // map[userID]remoteAddr

	service service.Service
	mu      sync.RWMutex
}

func (hub *websocketHub) RegisterWebsocket(w http.ResponseWriter, r *http.Request) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	tokenHeader := r.Header.Get("Authorization")
	slog.Debug("hub.RegisterWebsocket", "Authorization", tokenHeader)
	if tokenHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	bearerToken := strings.TrimPrefix(tokenHeader, "Bearer ")
	if bearerToken == tokenHeader {
		slog.Error("hub.RegisterWebsocket", "error", "mismatching tokenHeader and bearerToken")
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	id, err := hub.service.ValidateToken(r.Context(), bearerToken)
	if err != nil {
		slog.Error("hub.RegisterWebsocket failed to validate token", "error", err.Error())
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Failed to set websocket upgrade: ", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hub.clients[r.RemoteAddr] = models.WebsocketClient{
		Conn:   conn,
		UserID: id,
	}
	go hub.handleClientMessage(conn)
}

func (hub *websocketHub) CloseWebsocket(w http.ResponseWriter, r *http.Request) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	delete(hub.clients, r.RemoteAddr)
	hub.clients[r.RemoteAddr].Conn.Close()
}

// server sends message to specific client
func (hub *websocketHub) SendMessage(remoteAddr string, message []byte) error {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	client, ok := hub.clients[remoteAddr]
	if !ok {
		return ErrClientNotFound
	}

	if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

// server sends message to all clients
func (hub *websocketHub) broadcastMessage(message []byte, chatID, userID string) {
	clientIds, err := hub.service.GetChatMembers(context.Background(), chatID, userID)
	if err != nil {
		slog.Error("failed to get chat members", "error", err.Error())
		return
	}

	for _, id := range clientIds {
		if err := hub.SendMessage(id, message); err != nil {
			slog.Error("failed to send message to client", "error", err.Error())
		}
	}
}

// server receive message
func (hub *websocketHub) handleClientMessage(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("error reading message", "error", err)
			break
		}

		var userMsg dto.UserRequestDto
		if err := json.Unmarshal(message, &userMsg); err != nil {
			slog.Error("failed to unmarshal message", "error", err)
			continue
		}

		switch userMsg.Action {
		case enum.UserSendMessage:
			hub.userSendMessageController(conn, userMsg.Payload)
		case enum.UserGetMessages:
			hub.userAsksMessagesController(conn, userMsg.Payload)
		case enum.UserDeleteMessage:
			hub.userDeleteMessageController(conn, userMsg.Payload)
		}
	}
}

// controllers
func (hub *websocketHub) userSendMessageController(conn *websocket.Conn, payload []byte) {
	slog.Debug("userSendMessageController", "payload", string(payload), "conn", conn)
	var req dto.NewMessageDto
	if err := json.Unmarshal(payload, &req); err != nil {
		slog.Error("failed to unmarshal message", "error", err)
		return
	}
	msgpb := mapper.SaveMessageHttpRequestToPb(req)

	if err := hub.service.SaveMessage(context.Background(), msgpb); err != nil {
		slog.Error("failed to save message", "error", err)
		return
	}

	addr := conn.RemoteAddr().String()
	hub.broadcastMessage(payload, req.ChatID, hub.clients[addr].UserID)
}

func (hub *websocketHub) userDeleteMessageController(conn *websocket.Conn, payload []byte) {
	slog.Debug("userDeleteMessageController", "payload", string(payload), "conn", conn)
	var req dto.DeleteMessageDto
	if err := json.Unmarshal(payload, &req); err != nil {
		slog.Error("failed to unmarshal message", "error", err)
		return
	}

	if err := hub.service.DeleteMessage(context.Background(), req.ChatID, req.MessageID); err != nil {
		slog.Error("failed to delete message", "error", err)
		return
	}

	addr := conn.RemoteAddr().String()
	hub.broadcastMessage(payload, req.ChatID, hub.clients[addr].UserID)
}

func (hub *websocketHub) userAsksMessagesController(conn *websocket.Conn, payload []byte) {
	slog.Debug("userAsksMessagesController", "payload", string(payload), "conn", conn)
	var req dto.AskMessagesDto
	if err := json.Unmarshal(payload, &req); err != nil {
		slog.Error("failed to unmarshal message", "error", err)
		return
	}

	messages, err := hub.service.GetMessages(context.Background(), req.ChatID, req.Limit, req.Offset)
	if err != nil {
		slog.Error("failed to get messages", "error", err)
		return
	}

	rawResp, err := json.Marshal(messages)
	if err != nil {
		slog.Error("failed to marshal messages", "error", err)
		return
	}

	addr := conn.RemoteAddr().String()
	hub.broadcastMessage(rawResp, req.ChatID, hub.clients[addr].UserID)
}

func New(service service.Service) WebsocketHub {
	return &websocketHub{
		clients: make(map[string]models.WebsocketClient),
		service: service,
		mu:      sync.RWMutex{},
	}
}
