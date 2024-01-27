package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var port int = 8080
	router.GET("/ping", pingHandler)

	log.Printf("Starting server on port %v", port)
	router.Run(fmt.Sprintf(":%v", port))
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
