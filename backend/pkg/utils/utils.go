package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/arturshadnik/gobot/backend/internal/config"
	"github.com/arturshadnik/gobot/backend/internal/models"
)

func GetOpenAIResponse(requestConfig map[string]any, apiKey string) (string, error) {

	jsonData, err := json.Marshal(requestConfig)
	if err != nil {
		return "", err
	}
	reqBody := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", config.OpenAIUrl, reqBody)
	log.Print(reqBody)
	if err != nil {
		return "", err
	}

	if apiKey == "" {
		apiKey = config.OpenAIApiKey
		log.Print("using default key")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	aiMessage, err := parseJSON(body)
	if err != nil {
		return "", err
	}
	return aiMessage.Choices[0].Message.Content, nil
}

func parseJSON(jsonData []byte) (models.Completion, error) {
	var completion models.Completion

	err := json.Unmarshal(jsonData, &completion)

	if err != nil {
		return models.Completion{}, err
	}
	return completion, nil
}
