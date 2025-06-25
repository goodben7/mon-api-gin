package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/goodben7/mon-api-gin.git/models"
)

// @Summary Récupère un utilisateur par défaut
// @Description Retourne un utilisateur exemple (pour tests)
// @Tags users
// @Produce json
// @Success 200 {object} models.User
// @Router /api/user [get]
func GetUser(c *gin.Context) {
	user := models.User{
		ID:   "123",
		Name: "Jean Dupont",
	}
	c.JSON(200, user)
}

// @Summary Récupère un utilisateur par son ID
// @Description Retourne les détails d'un utilisateur spécifique
// @Tags users
// @Produce json
// @Param id path string true "ID de l'utilisateur"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string "Utilisateur non trouvé"
// @Router /api/user/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	// Simuler une DB
	user := models.User{
		ID:   id,
		Name: "Utilisateur " + id,
	}
	c.JSON(200, user)
}

// @Summary Crée un utilisateur
// @Description Crée un nouvel utilisateur avec validation des données
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body  models.User  true  "Infos utilisateur"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/user [post]
func CreateUser(c *gin.Context) {
	var user models.User

	// Lie et valide automatiquement
	if err := c.ShouldBindJSON(&user); err != nil {
		// Récupère les erreurs de validation
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErrors {
				// Messages d'erreur standards
				errorMessages = append(errorMessages,
					fmt.Sprintf("Erreur sur le champ '%s': %s",
						fieldError.Field(),
						fieldError.Tag()))
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
