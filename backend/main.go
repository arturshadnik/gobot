// @title           Gen AI capture the flag with Go
// @version         1.0
// @contact.name Artur Shad Nik
// @host gobot-backend-ew7jmfmn3a-uw.a.run.app

package main

import (
	_ "github.com/arturshadnik/gobot/backend/docs"
	"github.com/arturshadnik/gobot/backend/internal/api"
	"github.com/arturshadnik/gobot/backend/internal/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	db.InitFirestore()
	defer db.CloseFirestore()
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "https://goot.arturshadnik.com"}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"*"}

	router.Use(cors.New(config))

	router.GET("/ping", api.HealthCheck)
	router.POST("/chat/:id", api.BotResponse)
	router.GET("/chat/:id", api.GetMessages)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
