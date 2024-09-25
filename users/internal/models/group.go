package models

type Group struct {
	ID      string
	Name    string
	Avatar  *string
	Members []*User
}
