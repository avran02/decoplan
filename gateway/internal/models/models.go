package models

type UserChat struct {
	ID       string  `json:"id"`
	ChatName *string `json:"chatName,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}
