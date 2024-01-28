package api

import (
	"net/http"
	"time"

	"github.com/arturshadnik/gobot/backend/internal/models"
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
// @Param message body models.Message true "Message request"
// @Success 200
// @Router /chat [post]
func BotResponse(c *gin.Context) {
	var incomingMessage models.Message
	if err := c.BindJSON(&incomingMessage); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	outgoingMessage := models.Message{UUID: incomingMessage.UUID, Name: "gobot", Timestamp: time.Now(), Body: "hello"}

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

// for future reference
// HandleWithQuery godoc
// @Summary handle query
// @Description example for handling a request with a query
// @Tags chat
// @Accept  json
// @Produce  json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success 200
// @Router /query [get]
func HandlerWithQuery(c *gin.Context) {
	query := c.Query("query")
	c.JSON(http.StatusOK, gin.H{
		"query": query,
	})
}
