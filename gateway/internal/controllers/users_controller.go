package controllers

import (
	"fmt"
	"net/http"

	"github.com/avran02/decplan/gateway/internal/dto"
	"github.com/avran02/decplan/gateway/internal/enum"
	"github.com/avran02/decplan/gateway/internal/services"
	"github.com/go-chi/chi/v5"
)

type UsersController interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
}

type usersController struct {
	s services.UsersService
}

func (c *usersController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	if err := c.s.CreateUser(
		r.Context(),
		req.ID,
		req.Name,
		req.BirthDate,
	); err != nil {
		apiError(w, http.StatusInternalServerError, fmt.Errorf("failed to make gRPC call: %w", err))
		return
	}

	resp := dto.CreateUserResponse{
		Ok: true,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
	}
}

func (c *usersController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := c.s.GetUser(r.Context(), chi.URLParam(r, enum.UserID.String()))
	if err != nil {
		apiError(w, http.StatusInternalServerError, fmt.Errorf("failed to make gRPC call: %w", err))
		return
	}

	resp := dto.GetUserResponse{
		Name:      user.Name,
		Avatar:    user.Avatar,
		BirthDate: user.BirthDate.AsTime(),
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
	}
}

func (c *usersController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		apiError(w, http.StatusInternalServerError, err)
		return
	}

	if err := c.s.UpdateUser(
		r.Context(),
		chi.URLParam(r, enum.UserID.String()),
		req.Name,
		req.Avatar,
		req.BirthDate,
	); err != nil {
		apiError(w, http.StatusInternalServerError, fmt.Errorf("failed to make gRPC call: %w", err))
		return
	}

	resp := dto.UpdateUserResponse{
		Ok: true,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
	}
}

func (c *usersController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := c.s.DeleteUser(
		r.Context(),
		chi.URLParam(r, enum.UserID.String()),
	); err != nil {
		apiError(w, http.StatusInternalServerError, fmt.Errorf("failed to make gRPC call: %w", err))
		return
	}

	resp := dto.DeleteUserResponse{
		Ok: true,
	}
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		apiError(w, http.StatusInternalServerError, err)
	}
}

func newUsersController(srv services.UsersService) UsersController {
	return &usersController{
		s: srv,
	}
}
