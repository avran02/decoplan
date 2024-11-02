package dto

type DeleteMessageDto struct {
	MessageID uint64 `json:"messageId"`
	ChatID    string `json:"chatId"`
}
