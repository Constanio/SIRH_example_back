package main

import (
	"sirh/database"
	"sirh/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Chargement des variables d'environnement (.env)
	if err := godotenv.Load(); err != nil {
		log.Println("Note: Aucun fichier .env trouvé, utilisation des variables locales")
	}

	// Connexion à la base de données
	_, err := database.Connect()
	if err != nil {
		log.Fatalf("Échec de la connexion à la base de données: %v", err)
	}

	// Seeder
	database.Seed()

	// Initialisation de Gin
	r := gin.Default()

	// Configuration CORS
	config := cors.DefaultConfig()
	
	// Liste des origines autorisées (On ajoute les variantes avec et sans slash)
	config.AllowOrigins = []string{
		"http://localhost:5173",
		"https://sirhexamplefront.netlify.app",
		"https://sirhexamplefront.netlify.app/",
	}
	
	frontURL := os.Getenv("FRONTEND_URL")
	if frontURL != "" {
		config.AllowOrigins = append(config.AllowOrigins, frontURL)
	}
	
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// Configuration des routes
	routes.Setup(r)

	// Lancement du serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Serveur démarré sur le port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Impossible de lancer le serveur:", err)
	}
}
