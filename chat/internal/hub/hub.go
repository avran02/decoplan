package hub

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
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
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type WebsocketHub interface {
	CloseWebsocket(w http.ResponseWriter, r *http.Request)
	RegisterWebsocket(w http.ResponseWriter, r *http.Request)
}

type websocketHub struct {
	clientConnections map[string]models.WebsocketClient // map[remoteAddr]models.WebsocketClient
	clientIPs         map[string]string                 // map[userID]remoteAddr

	service service.Service
	mu      sync.RWMutex
}

func (hub *websocketHub) RegisterWebsocket(w http.ResponseWriter, r *http.Request) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	bearerToken := r.URL.Query().Get("token")
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
	hub.clientConnections[r.RemoteAddr] = models.WebsocketClient{
		Conn:   conn,
		UserID: id,
	}
	hub.clientIPs[id] = r.RemoteAddr

	go hub.handleClientMessage(conn)
}

func (hub *websocketHub) CloseWebsocket(w http.ResponseWriter, r *http.Request) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	delete(hub.clientConnections, r.RemoteAddr)
	hub.clientConnections[r.RemoteAddr].Conn.Close()
}

// server sends message to specific client
func (hub *websocketHub) sendMessage(remoteAddr string, message []byte) {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	client, ok := hub.clientConnections[remoteAddr]
	if !ok {
		slog.Error("client not found", "remoteAddr", remoteAddr)
		return
	}

	if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
		slog.Error("failed to send message", "error", err.Error())
	}
}

// server sends message to clients
func (hub *websocketHub) broadcastMessage(message []byte, chatID, userID string, skipIssuer bool) {
	slog.Info("hub.broadcastMessage")
	clientIds, err := hub.service.GetChatMembers(context.Background(), chatID, userID)
	if err != nil {
		slog.Error("failed to get chat members", "error", err.Error())
		return
	}

	for _, id := range clientIds {
		if id == userID && skipIssuer {
			continue
		}
		clientIP, exists := hub.clientIPs[id]
		if !exists {
			continue
		}
		hub.sendMessage(clientIP, message)
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
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		var userMsg dto.WSMessageDto
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
	slog.Info("userSendMessageController")
	var req dto.FromClientMessageDto
	if err := json.Unmarshal(payload, &req); err != nil {
		slog.Error("failed to unmarshal message", "error", err)
		return
	}

	// can be optimized by unmarshalling payload to pb
	msgpb := mapper.SaveMessageHttpRequestToPb(req, hub.clientConnections[conn.RemoteAddr().String()].UserID)
	msg, err := hub.service.SaveMessage(context.Background(), msgpb)
	if err != nil {
		slog.Error("failed to save message", "error", err)
		return
	}

	// add action to response
	resp, err := mapper.MessageToResponse(msg, enum.ServerSendMessage)
	if err != nil {
		slog.Error("failed to marshal message", "error", err)
		return
	}

	addr := conn.RemoteAddr().String()
	hub.broadcastMessage(resp, req.ChatID, hub.clientConnections[addr].UserID, true)
}

func (hub *websocketHub) userDeleteMessageController(conn *websocket.Conn, payload []byte) {
	slog.Debug("userDeleteMessageController", "payload", string(payload))
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
	hub.broadcastMessage(payload, req.ChatID, hub.clientConnections[addr].UserID, false)
}

func (hub *websocketHub) userAsksMessagesController(conn *websocket.Conn, payload []byte) {
	slog.Debug("userAsksMessagesController", "payload", string(payload))
	var req dto.UserAskMessagesDto
	if err := json.Unmarshal(payload, &req); err != nil {
		slog.Error("failed to unmarshal message", "error", err)
		return
	}

	messages, err := hub.service.GetMessages(context.Background(), req.ChatID, req.Limit, req.Offset)
	if err != nil {
		slog.Error("failed to get messages", "error", err)
		return
	}

	resp, err := mapper.MessagesToResponse(messages, enum.ServerSendMessage)
	if err != nil {
		slog.Error("failed to marshal message", "error", err)
		return
	}

	addr := conn.RemoteAddr().String()
	hub.sendMessage(addr, resp)
}

func New(service service.Service) WebsocketHub {
	return &websocketHub{
		clientConnections: make(map[string]models.WebsocketClient),
		clientIPs:         make(map[string]string),
		service:           service,
		mu:                sync.RWMutex{},
	}
}
