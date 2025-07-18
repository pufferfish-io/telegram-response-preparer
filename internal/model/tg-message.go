package model

type SendMessageRequest struct {
	ChatID              int64  `json:"chat_id"`
	Text                string `json:"text"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ReplyToMessageID    *int   `json:"reply_to_message_id,omitempty"`
}
