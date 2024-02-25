package models

import "time"

type ConvoMessage struct {
	MessageID string    `json:"uuid" example:"1"`
	Sender    string    `json:"name" example:"gobot"`
	Timestamp time.Time `json:"timestamp" example:"2022-01-01T00:00:00Z"`
	Body      string    `json:"body" example:"Hello my name is Artur"`
}

type Conversation struct {
	ConversationID string
	Messages       []string
	UserID         string
}
