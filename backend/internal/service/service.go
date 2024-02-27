// business logic layer

package service

import (
	"github.com/arturshadnik/gobot/backend/internal/config"
	"github.com/arturshadnik/gobot/backend/internal/db"
	"github.com/arturshadnik/gobot/backend/pkg/utils"
)

func ProcessIncomingMsg(message string, level string, id string) (string, error) {
	err := db.StoreMessage(message, level, "user", id)
	if err != nil {
		return "", err
	}

	conversation, err := db.LoadConversation(id, level)
	if err != nil {
		return "", err
	}

	messages, err := db.GetMessages(conversation.Messages)
	if err != nil {
		return "", err
	}

	convConfig, err := config.GetLevelSettings(level)
	if err != nil {
		return "", err
	}

	sysMsg := []map[string]string{{
		"role":    "system",
		"content": convConfig.Prompt,
	}}

	messages = append(sysMsg, messages...)

	openaiCompletionConfig := map[string]any{
		"model":       convConfig.Model,
		"temperature": 1,
		"messages":    messages,
	}

	resp, err := utils.GetOpenAIResponse(openaiCompletionConfig)
	if err != nil {
		return "", err
	}

	err = db.StoreMessage(resp, level, "assistant", id)
	if err != nil {
		return "", nil
	}

	return resp, nil
}
