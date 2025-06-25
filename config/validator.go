package config

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// CustomValidator est notre instance de validateur
var Validate *validator.Validate

func SetupValidator() {
	// Récupère le moteur de validation de Gin
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		Validate = v
	}
}
