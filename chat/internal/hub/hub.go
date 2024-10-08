package hub

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/avran02/decoplan/chat/enum"
	"github.com/avran02/decoplan/chat/internal/dto"
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

		var userMsg dto.UserRequest
		if err := json.Unmarshal(message, &userMsg); err != nil {
			slog.Error("failed to unmarshal message from %s: %v", remoteAddr, err)
			continue
		}

		switch userMsg.Action {
		case enum.UserSendMessage:
			hub.UserSendMessageController(conn, userMsg.Payload)
		case enum.UserGetMessages:
			hub.UserAsksMessagesController(conn, userMsg.Payload)
		case enum.UserDeleteMessage:
			hub.UserDeleteMessageController(conn, userMsg.Payload)
		}
	}
}

// controllers
func (hub *websocketHub) UserSendMessageController(conn *websocket.Conn, payload []byte) {}

func (hub *websocketHub) UserDeleteMessageController(conn *websocket.Conn, payload []byte) {}

func (hub *websocketHub) UserAsksMessagesController(conn *websocket.Conn, payload []byte) {}

func New() WebsocketHub {
	return &websocketHub{
		clients: make(map[string]*websocket.Conn),
		mu:      sync.RWMutex{},
	}
}
