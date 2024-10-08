package dto

import "time"

type NewMessage struct {
	Sender    string    `json:"sender"`
	Content   Content   `json:"content"`
	TimeStamp time.Time `json:"timestamp"`
}

type Content struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	URL string `json:"url"`
}
