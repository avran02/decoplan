package mapper

import (
	"log/slog"

	"github.com/avran02/decoplan/chat/enum"
	"github.com/avran02/decoplan/chat/internal/dto"
	"github.com/avran02/decoplan/chat/internal/models"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func MessageToResponse(msg *models.Message, act enum.MessageTypes) ([]byte, error) {
	messages := []models.Message{*msg}
	payload, err := json.Marshal(messages)
	if err != nil {
		slog.Error("failed to marshal message", "error", err.Error())
		return nil, err
	}

	res := dto.WSMessageDto{
		Action:  act,
		Payload: payload,
	}

	return json.Marshal(res)
}

func MessagesToResponse(messages []models.Message, act enum.MessageTypes) ([]byte, error) {
	payload, err := json.Marshal(messages)
	if err != nil {
		slog.Error("failed to marshal message", "error", err.Error())
		return nil, err
	}

	res := dto.WSMessageDto{
		Action:  act,
		Payload: payload,
	}

	return json.Marshal(res)
}
