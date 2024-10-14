package models

import (
	"encoding/json"
	"time"
)

type Message struct {
	ID     uint64 `json:"id"`
	ChatID string `json:"chatId"`

	Sender  string `json:"sender"`
	Content string `json:"content"`

	Attachments []Attachment `json:"attachments"`

	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
func (m *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}

type Attachment struct {
	ID        string `json:"id"`
	MessageID uint64 `json:"messageId"`
	ChatID    string `json:"chatId"`
	URL       string `json:"url"`
}

func (a Attachment) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Attachment) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &a)
}
