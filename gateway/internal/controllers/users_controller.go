package controllers

import "net/http"

type UsersController interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	UpdateUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
}

type usersController struct{}

func (c *usersController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {}
func (c *usersController) GetUserHandler(w http.ResponseWriter, r *http.Request)    {}
func (c *usersController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {}
func (c *usersController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {}

func NewUsersController() UsersController {
	return &usersController{}
}
