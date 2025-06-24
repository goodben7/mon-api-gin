package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/goodben7/mon-api-gin.git/models"
)

func GetUser(c *gin.Context) {
	user := models.User{
		ID:   "123",
		Name: "Jean Dupont",
	}
	c.JSON(200, user)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	// Simuler une DB
	user := models.User{
		ID:   id,
		Name: "Utilisateur " + id,
	}
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	// Lie et valide automatiquement
	if err := c.ShouldBindJSON(&user); err != nil {
		// Récupère les erreurs de validation
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErrors {
				// Message d'erreur personnalisé pour le tag "not_admin"
				if fieldError.Tag() == "not_admin" {
					errorMessages = append(errorMessages, "Le nom 'admin' est interdit")
				} else {
					// Message d'erreur par défaut pour les autres tags
					errorMessages = append(errorMessages,
						fmt.Sprintf("Erreur sur le champ '%s': %s",
							fieldError.Field(),
							fieldError.Tag()))
				}
			}
			c.JSON(400, gin.H{"errors": errorMessages})
			return
		}
		// Autres types d'erreurs (JSON malformé, etc.)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Si tout est valide
	c.JSON(201, gin.H{"status": "user created", "data": user})
}
