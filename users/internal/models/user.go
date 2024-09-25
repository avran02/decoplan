package models

import "time"

type User struct {
	ID        string
	Name      string
	BirthDate time.Time
	Avatar    *string
}

type UpdateUser struct {
	ID        string
	Name      *string
	BirthDate *time.Time
	Avatar    *string
}
