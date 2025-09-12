package contract

import "time"

type NormalizedResponse struct {
	ChatID         int64     `json:"chat_id"`
	Text           string    `json:"text,omitempty"`
	Silent         bool      `json:"silent,omitempty"`
	Context        any       `json:"context,omitempty"`
	Source         string    `json:"source"`
	UserID         int64     `json:"user_id"`
	Username       *string   `json:"username,omitempty"`
	Timestamp      time.Time `json:"timestamp"`
	OriginalUpdate any       `json:"original_update"`
}

type SendMessageRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
