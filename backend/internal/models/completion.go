package models

import (
	"encoding/json"
)

type Completion struct {
	ID      string
	Object  string
	Created int
	Model   string
	Usage   Usage
	Choices []Choice
}

type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

type Choice struct {
	Message      Message
	LogProbs     *json.RawMessage
	FinishReason string
	Index        int
}

type Message struct {
	Role    string
	Content string
}
