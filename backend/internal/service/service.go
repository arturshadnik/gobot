// business logic layer

package service

import (
	"time"

	"github.com/arturshadnik/gobot/backend/internal/config"
	"github.com/arturshadnik/gobot/backend/internal/db"
	"github.com/arturshadnik/gobot/backend/internal/models"
	"github.com/arturshadnik/gobot/backend/pkg/utils"
)

func ProcessIncomingMsg(message string, level string, id string) (models.ConvoMessage, error) {
	err := db.StoreMessage(message, level, "user", id)
	if err != nil {
		return models.ConvoMessage{}, err
	}

	conversation, err := db.LoadConversation(id, level)
	if err != nil {
		return models.ConvoMessage{}, err
	}

	messages, err := db.GetMessages(conversation.Messages)
	if err != nil {
		return models.ConvoMessage{}, err
	}

	for _, msg := range messages {
		delete(msg, "timestamp")
	}

	convConfig, err := config.GetLevelSettings(level)
	if err != nil {
		return models.ConvoMessage{}, err
	}

	sysMsg := []map[string]any{{
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
		return models.ConvoMessage{}, err
	}

	err = db.StoreMessage(resp, level, "assistant", id)
	if err != nil {
		return models.ConvoMessage{}, nil
	}

	return models.ConvoMessage{Role: "assistant", Content: resp, Timestamp: time.Now()}, nil
}
