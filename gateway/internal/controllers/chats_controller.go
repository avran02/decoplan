package controllers

import (
	"net/http"

	"github.com/avran02/decplan/gateway/internal/dto"
	"github.com/avran02/decplan/gateway/internal/enum"
	"github.com/avran02/decplan/gateway/internal/mapper"
	"github.com/avran02/decplan/gateway/internal/services"
	"github.com/go-chi/chi/v5"
)

type ChatsController interface {
	CreateChatHandler(w http.ResponseWriter, r *http.Request)
	GetChatHandler(w http.ResponseWriter, r *http.Request)
	DeleteChatHandler(w http.ResponseWriter, r *http.Request)

	RemoveUserFromChatHandler(w http.ResponseWriter, r *http.Request)
	AddUserToChatHandler(w http.ResponseWriter, r *http.Request)
}

type chatsController struct {
	s services.UsersService
}

func (c *chatsController) CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateChatRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	chatID, err := c.s.CreateChat(r.Context(), req.Name, req.UserIDs)
	if err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	resp := dto.CreateChatResponse{
		ChatID: chatID,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
	}
}

func (c *chatsController) GetChatHandler(w http.ResponseWriter, r *http.Request) {
	chatInfo, err := c.s.GetChat(r.Context(), chi.URLParam(r, enum.ChatID.String()))
	if err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	resp := mapper.ChatInfoFromPbToHttp(chatInfo)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}
}

func (c *chatsController) DeleteChatHandler(w http.ResponseWriter, r *http.Request) {
	if err := c.s.DeleteChat(r.Context(), chi.URLParam(r, enum.ChatID.String())); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	resp := dto.DeleteChatResponse{
		Ok: true,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
	}
}

func (c *chatsController) RemoveUserFromChatHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.RemoveUserFromChatRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	if err := c.s.RemoveUserFromChat(
		r.Context(),
		req.UserID,
		chi.URLParam(r, enum.ChatID.String())); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	resp := dto.RemoveUserFromChatResponse{
		Ok: true,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
	}
}

func (c *chatsController) AddUserToChatHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.AddUserToChatRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	if err := c.s.AddUserToChat(r.Context(),
		chi.URLParam(r, enum.ChatID.String()),
		req.UserID); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	resp := dto.AddUserToChatResponse{
		Ok: true,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
	}
}

func NewChatsController() ChatsController {
	return &chatsController{}
}
