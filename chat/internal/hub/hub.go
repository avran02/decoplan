package hub

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/avran02/decoplan/chat/enum"
	"github.com/avran02/decoplan/chat/internal/dto"
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
	clients map[string]*websocket.Conn
	service service.Service
	mu      sync.RWMutex
}

func (hub *websocketHub) RegisterWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Failed to set websocket upgrade: ", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hub.clients[r.RemoteAddr] = conn
	go hub.handleClientMessage(r.RemoteAddr, conn)
}

func (hub *websocketHub) CloseWebsocket(w http.ResponseWriter, r *http.Request) {
	if err := hub.clients[r.RemoteAddr].Close(); err != nil {
		slog.Error("Failed to close websocket: ", "error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	delete(hub.clients, r.RemoteAddr)
}

// server sends message to specific client
func (hub *websocketHub) SendMessage(remoteAddr string, message []byte) error {
	hub.mu.RLock()
	defer hub.mu.RUnlock()
	conn, ok := hub.clients[remoteAddr]
	if !ok {
		return ErrClientNotFound
	}
	if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

// server sends message to all clients
func (hub *websocketHub) broadcastMessage(message []byte) {
	for addr := range hub.clients {
		if err := hub.SendMessage(addr, message); err != nil {
			slog.Error("failed to send message to client", "error", err.Error())
		}
	}
}

// server receive message
func (hub *websocketHub) handleClientMessage(remoteAddr string, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("error reading message from %s: %v", remoteAddr, err)
			break
		}

		var userMsg dto.UserRequestDto
		if err := json.Unmarshal(message, &userMsg); err != nil {
			slog.Error("failed to unmarshal message from %s: %v", remoteAddr, err)
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

	if err := hub.service.SaveMessage(context.Background(), req.Message); err != nil {
		slog.Error("failed to save message", "error", err)
		return
	}
	hub.broadcastMessage(payload)
}

func (hub *websocketHub) userDeleteMessageController(conn *websocket.Conn, payload []byte) {
	slog.Debug("userDeleteMessageController", "payload", string(payload), "conn", conn)
	var req dto.DeleteMessageDto
	if err := json.Unmarshal(payload, &req); err != nil {
		slog.Error("failed to unmarshal message", "error", err)
		return
	}
}

func (hub *websocketHub) userAsksMessagesController(conn *websocket.Conn, payload []byte) {
	slog.Debug("userAsksMessagesController", "payload", string(payload), "conn", conn)
	var req dto.AskMessagesDto
	if err := json.Unmarshal(payload, &req); err != nil {
		slog.Error("failed to unmarshal message", "error", err)
		return
	}
}

func New(service service.Service) WebsocketHub {
	return &websocketHub{
		clients: make(map[string]*websocket.Conn),
		service: service,
		mu:      sync.RWMutex{},
	}
}
