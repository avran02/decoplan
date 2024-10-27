package dto

type CreateChatRequest struct {
	Name    string   `json:"name"`
	UserIDs []string `json:"userIDs"`
}

type CreateChatResponse struct {
	ChatID string `json:"chatID"`
}

type GetChatResponse struct {
	ID       string       `json:"id"`
	ChatName *string      `json:"chatName,omitempty"`
	Avatar   *string      `json:"avatar,omitempty"`
	Members  []UserMember `json:"members"`
}

type UserMember struct {
	UserID string `json:"userID"`
}

type AddUserToChatRequest struct {
	UserID string `json:"userID"`
}

type AddUserToChatResponse struct {
	Ok bool `json:"ok"`
}

type RemoveUserFromChatRequest struct {
	UserID string `json:"userID"`
}

type RemoveUserFromChatResponse struct {
	Ok bool `json:"ok"`
}

type DeleteChatRequest struct {
	ID string `json:"id"`
}

type DeleteChatResponse struct {
	Ok bool `json:"ok"`
}
