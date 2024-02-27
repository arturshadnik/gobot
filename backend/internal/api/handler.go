package api

import (
	"log"
	"net/http"

	"github.com/arturshadnik/gobot/backend/internal/db"
	"github.com/arturshadnik/gobot/backend/internal/service"
	"github.com/gin-gonic/gin"
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
// @Success 200
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
			c.JSON(http.StatusOK, "Conversation History Reset.")
		}
		return
	}
	outgoingMessage, err := service.ProcessIncomingMsg(query, level, id)
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
	convo, err := db.LoadConversation(id, level)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusNotFound, gin.H{"detail": "Conversation not found"})
		return
	}

	messages, err := db.GetMessages(convo.Messages)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
