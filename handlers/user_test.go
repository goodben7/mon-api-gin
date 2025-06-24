package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goodben7/mon-api-gin.git/handlers"
	"github.com/goodben7/mon-api-gin.git/models"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/user", handlers.CreateUser)
	return r
}

func TestCreateUser_ValidData(t *testing.T) {
	router := setupRouter()

	user := models.User{
		ID:    "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		Name:  "Valid User",
		Email: "valid@example.com",
		Age:   25,
	}
	body, _ := json.Marshal(user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "user created")
}

func TestCreateUser_InvalidData(t *testing.T) {
	tests := []struct {
		name     string
		payload  map[string]interface{}
		expected int      // Code HTTP attendu
		errors   []string // Messages d'erreur attendus
	}{
		{
			name: "Email invalide",
			payload: map[string]interface{}{
				"id":    "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"name":  "Test",
				"email": "invalid",
				"age":   25,
			},
			expected: http.StatusBadRequest,
			errors:   []string{"email"},
		},
		{
			name: "Age trop jeune",
			payload: map[string]interface{}{
				"id":    "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"name":  "Test",
				"email": "valid@example.com",
				"age":   17,
			},
			expected: http.StatusBadRequest,
			errors:   []string{"gte"},
		},
		{
			name: "Nom trop court",
			payload: map[string]interface{}{
				"id":    "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
				"name":  "A",
				"email": "valid@example.com",
				"age":   25,
			},
			expected: http.StatusBadRequest,
			errors:   []string{"min"},
		},
	}

	router := setupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expected, w.Code)

			// Vérifie que toutes les erreurs attendues sont présentes
			responseBody := w.Body.String()
			for _, err := range tt.errors {
				assert.Contains(t, responseBody, err)
			}
		})
	}
}
