package dto

import (
	"encoding/json"

	"github.com/avran02/decoplan/chat/enum"
)

type UserRequestDto struct {
	Action  enum.UserMessages `json:"act"`
	Payload json.RawMessage   `json:"payload"`
}
