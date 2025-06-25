package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// DevLogger configure un logger complet pour le développement
func DevLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Enregistre le temps de début
		start := time.Now()

		// Capture le corps de la requête pour les POST/PUT/PATCH
		var requestBody bytes.Buffer
		if c.Request.Body != nil && (c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH") {
			c.Request.Body = teeBody(c.Request.Body, &requestBody)
		}

		// Traitement de la requête
		c.Next()

		// Préparation des données de log
		logData := map[string]interface{}{
			"time":      time.Now().Format("2006-01-02 15:04:05"),
			"status":    c.Writer.Status(),
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"duration":  time.Since(start).String(),
			"clientIP":  c.ClientIP(),
			"userAgent": c.Request.UserAgent(),
		}

		// Ajoute le corps de la requête si pertinent
		if requestBody.Len() > 0 {
			var body interface{}
			if err := json.Unmarshal(requestBody.Bytes(), &body); err == nil {
				logData["body"] = body
			} else {
				logData["body"] = requestBody.String()
			}
		}

		// Formatage et affichage du log
		logJSON, _ := json.Marshal(logData)
		log.Printf("[DEV] %s\n", string(logJSON))
	}
}

// teeBody permet de lire le corps sans le consommer
func teeBody(src io.ReadCloser, dst *bytes.Buffer) io.ReadCloser {
	tee := io.TeeReader(src, dst)
	return io.NopCloser(tee)
}
