// @title           AI chatbot with Go
// @version         1.0
// @contact.name Artur Shad Nik
// @host localhost:8080

package main

import (
	_ "github.com/arturshadnik/gobot/backend/docs"
	"github.com/arturshadnik/gobot/backend/internal/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()

	router.GET("/ping", api.HealthCheck)
	router.POST("/chat", api.BotResponse)
	router.GET("/messages/:id", api.GetMessages)
	router.GET("/query", api.HandlerWithQuery)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
