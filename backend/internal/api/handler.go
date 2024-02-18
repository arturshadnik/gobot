package api

import (
	"log"
	"net/http"

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
// @Success 200
// @Router /chat [post]
func BotResponse(c *gin.Context) {

	query := c.Query("message")

	outgoingMessage, err := service.ProcessIncomingMsg(query)
	if err != nil {
		log.Printf("Error processing message %v ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, outgoingMessage)
}

// for future reference
// GetMessages godoc
// @Summary Get messages
// @Description handler to fetch all past messages between the user and the bot
// @Tags chat
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "Account ID"
// @Success 200
// @Router /messages/{id} [get]
func GetMessages(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
