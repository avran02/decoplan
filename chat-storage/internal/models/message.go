package models

import (
	"encoding/json"
	"time"
)

type Message struct {
	ID     uint64 `bson:"_id"`
	ChatID string

	Sender  string
	Content string

	Attachments []Attachment

	CreatedAt time.Time
	DeletedAt *time.Time
}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
func (m *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &m)
}

type Attachment struct {
	ID        string `bson:"_id"`
	MessageID uint64
	ChatID    string
	URL       string
}

func (a Attachment) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Attachment) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &a)
}
