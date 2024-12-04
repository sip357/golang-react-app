// Routes for authenticated user
package routes

import (
	auth "go-project/go-backend/handlers/authOps"
	"net/http"

	"github.com/go-chi/chi/v5" // Chi Router package
	"github.com/rs/cors"       // CORS middleware package
)

// UserRoutes with CORS middleware
func AuthRoutes() http.Handler {
	// Create a new Chi router
	r := chi.NewRouter()

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000"}, // Allow frontend from localhost:3000
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true, // Allow credentials (cookies),
		AllowedHeaders:   []string{"Content-Type"},
	})

	// Apply the CORS middleware to the router
	r.Use(c.Handler)

	// Define routes with their respective handlers
	r.Post("/v1/create", auth.AuthCreateTask)
	r.Post("/auth", auth.ProtectedRoute)

	return r
}
