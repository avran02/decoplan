package dto

import "time"

type NewMessageDto struct {
	Sender    string    `json:"sender"`
	ChatID    string    `json:"chatId"`
	Content   Content   `json:"content"`
	TimeStamp time.Time `json:"timestamp"`
}

type Content struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
