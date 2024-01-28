package models

import "time"

type Message struct {
	UUID      int       `json:"uuid" example:"1"`
	Name      string    `json:"name" example:"gobot"`
	Timestamp time.Time `json:"timestamp" example:"2022-01-01T00:00:00Z"`
	Body      string    `json:"body" example:"Hello my name is Artur"`
}
