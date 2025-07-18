package service

import (
	"tg-response-preparer/internal/model"
)

type MessageConvertor struct {
}

func NewMessageConvertor() *MessageConvertor {
	return &MessageConvertor{}
}

func (MessageConvertor) RequestToTelegramMessage(input model.NormalizedResponse) *model.SendMessageRequest {
	return &model.SendMessageRequest{
		ChatID:              input.ChatID,
		Text:                *input.Text,
		DisableNotification: input.Silent,
		ReplyToMessageID:    input.ReplyToMessageID,
	}
}
