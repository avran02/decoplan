package controllers

import (
	"github.com/avran02/decplan/gateway/internal/services"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Controller struct {
	uc UsersController
	cc ChatsController
}

func (c Controller) UsersController() UsersController {
	return c.uc
}

func (c Controller) ChatsController() ChatsController {
	return c.cc
}

func New(srv services.UsersService) *Controller {
	return &Controller{
		uc: newUsersController(srv),
		cc: newChatsController(srv),
	}
}
