package enum

type URLParam string

func (param URLParam) String() string {
	return string(param)
}

const (
	ChatID = URLParam("chatID")
	UserID = URLParam("userID")
)
