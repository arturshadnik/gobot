// @title           AI chatbot with Go
// @version         1.0
// @contact.name Artur Shad Nik
// @host localhost:8080

package main

import (
	_ "github.com/arturshadnik/gobot/backend/docs"
	"github.com/arturshadnik/gobot/backend/internal/api"
	"github.com/arturshadnik/gobot/backend/internal/db"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	db.InitFirestore()
	defer db.CloseFirestore()
	router := gin.Default()

	router.GET("/ping", api.HealthCheck)
	router.POST("/chat/:id", api.BotResponse)
	router.GET("/chat/:id", api.GetMessages)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
