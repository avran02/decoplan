package enum

type UserMessages uint

const (
	UserGetMessages UserMessages = iota
	UserSendMessage
	UserDeleteMessage
)
