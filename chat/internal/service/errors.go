package service

import (
	"errors"
)

var (
	ErrUnknownError  = errors.New("unknown error")
	ErrChatForbidden = errors.New("forbidden chat")
)
