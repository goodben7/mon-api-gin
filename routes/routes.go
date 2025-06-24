package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/goodben7/mon-api-gin.git/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/user/:id", handlers.GetUserByID)
		api.POST("/user", handlers.CreateUser)
	}
}
