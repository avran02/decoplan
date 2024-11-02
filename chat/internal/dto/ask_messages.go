package dto

type AskMessagesDto struct {
	ChatID string `json:"chatId"`
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}
