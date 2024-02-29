package models

import "time"

type ConvoMessage struct {
	Role      string
	Content   string
	Timestamp time.Time
}

type Conversation struct {
	Messages []string
}
