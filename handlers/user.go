package handlers

import "github.com/gin-gonic/gin"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetUser(c *gin.Context) {
	user := User{
		ID:   "123",
		Name: "Jean Dupont",
	}
	c.JSON(200, user)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	// Simuler une DB
	user := User{
		ID:   id,
		Name: "Utilisateur " + id,
	}
	c.JSON(200, user)
}
