package models

type Chat struct {
	ID      string
	Name    string
	Avatar  *string
	Members []*User
}
