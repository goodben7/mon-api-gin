package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/goodben7/mon-api-gin.git/config"
	"github.com/goodben7/mon-api-gin.git/middleware"
	"github.com/goodben7/mon-api-gin.git/routes"
)

func main() {
	// Mode release (commenter pour le mode debug)
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	config.SetupValidator()

	routes.SetupRoutes(r)

	r.Use(middleware.Logger())

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
