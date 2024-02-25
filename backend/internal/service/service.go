// business logic layer

package service

import (
	"github.com/arturshadnik/gobot/backend/internal/db"
	"github.com/arturshadnik/gobot/backend/pkg/utils"
)

func ProcessIncomingMsg(message string, level string, id string) (string, error) {

	err := db.StoreMessage(message, level, "user")
	if err != nil {
		return "", err
	}
	var messages []map[string]string

	sysMsg := map[string]string{
		"role":    "system",
		"content": "You are an experienced coding mentor. Help the user with their questions",
	}

	userMsg := map[string]string{
		"role": "user", "content": message,
	}
	messages = append(messages, sysMsg, userMsg)

	openaiCompletionConfig := map[string]any{
		"model":       "gpt-3.5-turbo",
		"temperature": 0.3,
		"messages":    messages,
	}

	resp, err := utils.GetOpenAIResponse(openaiCompletionConfig)

	if err != nil {
		return "", err
	}

	err = db.StoreMessage(resp, level, "assistant")

	if err != nil {
		return "", nil
	}

	return resp, nil
}
