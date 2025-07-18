package model

type NormalizedResponse struct {
	ChatID           int64   `json:"chat_id"`
	Text             *string `json:"text,omitempty"`
	Silent           bool    `json:"silent,omitempty"`
	ReplyToMessageID *int    `json:"reply_to_message_id,omitempty"`
	Context          any     `json:"context,omitempty"`
}
