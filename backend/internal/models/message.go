package models

import "time"

type ConvoMessage struct {
	Role      string
	Content   string
	timestamp time.Time
}

type Conversation struct {
	Messages []string
}
