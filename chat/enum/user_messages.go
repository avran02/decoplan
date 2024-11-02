package enum

type MessageTypes uint

const (
	UserGetMessages MessageTypes = iota
	UserSendMessage
	UserDeleteMessage

	ServerSendMessage
	ServerDeleteMessage
)
