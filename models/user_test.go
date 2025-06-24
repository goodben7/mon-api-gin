package models_test

import (
	"testing"

	"github.com/goodben7/mon-api-gin.git/models"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		user    models.User
		isValid bool
	}{
		{
			name: "Utilisateur valide",
			user: models.User{
				ID:    "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				Name:  "Valid User",
				Email: "valid@example.com",
				Age:   25,
			},
			isValid: true,
		},
		{
			name: "Email invalide",
			user: models.User{
				ID:    "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				Name:  "Test",
				Email: "invalid-email",
				Age:   25,
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.user)
			assert.Equal(t, tt.isValid, err == nil)
		})
	}
}
