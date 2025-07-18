package processor

import (
	"context"
	"encoding/json"

	"tg-response-preparer/internal/messaging"
	"tg-response-preparer/internal/model"
	"tg-response-preparer/internal/service"
)

type TgMessagePreparer struct {
	convertor  *service.MessageConvertor
	kafkaTopic string
}

func NewTgMessagePreparer(kafkaTopic string, convertor *service.MessageConvertor) *TgMessagePreparer {
	return &TgMessagePreparer{
		convertor:  convertor,
		kafkaTopic: kafkaTopic,
	}
}

func (t *TgMessagePreparer) Handle(ctx context.Context, raw []byte) error {
	var requestMessage model.NormalizedResponse
	if err := json.Unmarshal(raw, &requestMessage); err != nil {
		return err
	}

	var msg = t.convertor.RequestToTelegramMessage(requestMessage)

	out, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return messaging.Send(t.kafkaTopic, out)
}
