package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/arturshadnik/gobot/backend/internal/db"
	"github.com/arturshadnik/gobot/backend/internal/models"
	"github.com/arturshadnik/gobot/backend/internal/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// BotResponse godoc
// @Summary Chat with a bot
// @Description Get a response to a question
// @Tags chat
// @Accept  json
// @Produce  json
// @Param message query string true "Message"
// @Param level query string true "Difficulty Level"
// @Param id path int true "Account ID"
// @Success 200 {object} models.ConvoMessage
// @Failure 500
// @Router /chat/{id} [post]
func BotResponse(c *gin.Context) {

	query := c.Query("message")
	level := c.Query("level")
	id := c.Param("id")

	if query == "conversation:reset" {
		err := db.ClearConversation(level + id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error"})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
		return
	}
	var apiKey string
	if len(body) > 0 {
		var reqBody models.Request
		if err := json.Unmarshal(body, &reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		apiKey = reqBody.ApiKey
	}

	outgoingMessage, err := service.ProcessIncomingMsg(query, level, id, apiKey)
	if err != nil {
		log.Printf("Error processing message %v ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, outgoingMessage)
}

// GetMessages godoc
// @Summary Get messages
// @Description handler to fetch all past messages between the user and the bot
// @Tags chat
// @Accept  json
// @Produce  json
// @Param level query string true "Difficulty Level"
// @Param        id   path      int  true  "Account ID"
// @Success 200
// @Router /chat/{id} [get]
func GetMessages(c *gin.Context) {

	level := c.Query("level")
	id := c.Param("id")
	messages, err := db.GetMessages(level + id)
	if err != nil {
		if status.Code(err) == codes.NotFound || len(messages) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"detail": "Conversation Not Found"})
			return
		}
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
