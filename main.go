package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Mode release (commenter pour le mode debug)
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Middleware global (ex: CORS)
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})

	// Route test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Démarrer le serveur
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Impossible de démarrer le serveur:", err)
	}
}