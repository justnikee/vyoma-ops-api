package main

import (
	"log"
	"net/http"
	"vyoma-api/internal/db"

	"vyoma-api/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"

	"github.com/joho/godotenv"
)

func main() {
	// loadign .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// connecting to database
	db.Connect()
	db.Ping()
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	r.Use(c.Handler)

	routes.RegisterRoutes(r)
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
