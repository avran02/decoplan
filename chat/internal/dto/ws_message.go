package dto

import (
	"encoding/json"

	"github.com/avran02/decoplan/chat/enum"
)

type WSMessageDto struct {
	Action  enum.MessageTypes `json:"act"`
	Payload json.RawMessage   `json:"payload"`
}
