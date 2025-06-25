package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/goodben7/mon-api-gin.git/docs" // docs générées par swag
	"github.com/goodben7/mon-api-gin.git/middleware"
	"github.com/goodben7/mon-api-gin.git/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mon API Gin - DEV
// @version 1.0-dev
// @description Documentation de l'API en développement
// @host localhost:8080
// @BasePath /
func main() {
	// 1. Configuration de l'environnement
	gin.SetMode(gin.DebugMode)

	// 2. Initialisation du routeur avec gin.New() au lieu de gin.Default()
	// (gin.Default() inclut déjà des middlewares de base qu'on veut personnaliser)
	r := gin.New()

	// 3. Middlewares de base RECOMMANDÉS pour tous les environnements
	r.Use(
		gin.Recovery(),         // Middleware de récupération après panic
		middleware.DevLogger(), // Notre logger personnalisé amélioré
	)

	// 4. Middleware CORS spécifique au dev
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// 5. Swagger (uniquement en dev)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1), // Désactive les models dans Swagger UI
	))

	// 6. Initialisation des routes
	routes.SetupRoutes(r)

	// 7. Route de santé enrichie
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"time":    time.Now().Format(time.RFC3339Nano),
			"status":  "healthy",
			"mode":    "development",
			"endpoints": []string{
				"/swagger/index.html",
				"/api/user",
				// Ajoutez d'autres endpoints ici
			},
		})
	})

	// 8. Message de démarrage amélioré
	startupMessage := `
	🚀 API Gin en mode développement
	📚 Swagger UI: http://localhost:8080/swagger/index.html
	🩺 Health check: http://localhost:8080/ping
	📝 Logs: Actifs avec format structuré
	`
	log.Println(startupMessage)

	// 9. Lancement du serveur avec gestion d'erreur
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Impossible de démarrer le serveur: %v", err)
	}
}
