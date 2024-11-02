package dto

import "github.com/avran02/decoplan/chat/internal/models"

type FromClientMessageDto struct {
	ChatID  string  `json:"chatId"`
	Content Content `json:"content"`
}

type Content struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type FromServerMessageDto []models.Message
