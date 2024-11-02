package dto

import "time"

type CreateUserRequest struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
}

type CreateUserResponse struct {
	Ok bool `json:"ok"`
}

type GetUserResponse struct {
	Name      string    `json:"name"`
	Avatar    *string   `json:"avatar,omitempty"`
	BirthDate time.Time `json:"birthDate"`
}

type UpdateUserRequest struct {
	Name      *string    `json:"name,omitempty"`
	Avatar    *string    `json:"avatar,omitempty"`
	BirthDate *time.Time `json:"birthDate,omitempty"`
}

type UpdateUserResponse struct {
	Ok bool `json:"ok"`
}

type DeleteUserResponse struct {
	Ok bool `json:"ok"`
}
