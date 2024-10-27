package controllers

import "net/http"

type ChatsController interface {
	CreateChatHandler(w http.ResponseWriter, r *http.Request)
	GetChatHandler(w http.ResponseWriter, r *http.Request)
	DeleteChatHandler(w http.ResponseWriter, r *http.Request)

	RemoveUserFromChatHandler(w http.ResponseWriter, r *http.Request)
	AddUserToChatHandler(w http.ResponseWriter, r *http.Request)
}

type chatsController struct{}

func (c *chatsController) CreateChatHandler(w http.ResponseWriter, r *http.Request)         {}
func (c *chatsController) GetChatHandler(w http.ResponseWriter, r *http.Request)            {}
func (c *chatsController) DeleteChatHandler(w http.ResponseWriter, r *http.Request)         {}
func (c *chatsController) RemoveUserFromChatHandler(w http.ResponseWriter, r *http.Request) {}
func (c *chatsController) AddUserToChatHandler(w http.ResponseWriter, r *http.Request)      {}

func NewChatsController() ChatsController {
	return &chatsController{}
}
