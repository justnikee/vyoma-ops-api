package routes

import (
	"net/http"

	"vyoma-api/internal/handlers"
	"vyoma-api/internal/repositories"
	"vyoma-api/internal/services"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	// deps

	businessRepo := &repositories.BusinessRepository{}
	businessService := services.NewBusinessService(businessRepo)
	businessHandler := handlers.NewBusinessHandler(businessService)

	r.Route("/businesses", func(r chi.Router) {
		r.Post("/", businessHandler.CreateBusiness)
		r.Get("/{id}", businessHandler.GetBusinessByID)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Get("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go backend 👋"))
	})
}
