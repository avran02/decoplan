package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	// Извлекаем URL для WebSocket-соединения из ответа
	wsURL := "ws://localhost:3002/connect"

	// Устанавливаем соединение WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()

	// Запуск горутины для чтения входящих сообщений
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			fmt.Printf("Received: %s\n", message)
		}
	}()

	// Основной цикл для отправки сообщений
	fmt.Println("Enter messages to send to the server. Type 'exit' to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}

		// Отправляем сообщение на сервер
		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading from input: %v", err)
	}
}
