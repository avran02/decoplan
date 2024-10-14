package models

import "github.com/gorilla/websocket"

type WebsocketClient struct {
	UserID string
	Conn   *websocket.Conn
}
