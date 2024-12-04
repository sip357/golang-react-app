package routes

import (
	handlers "go-project/go-backend/handlers/tasks" // import handlers

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors" // CORS middleware package
)

// TaskRoutes with CORS middleware
func TaskRoutes() http.Handler {
	r := chi.NewRouter()

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000"}, // Allow frontend from localhost:3000
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true, // Allow credentials (cookies),
		AllowedHeaders:   []string{"Content-Type"},
	})

	r.Use(c.Handler)

	// Routes associated with handlers
	r.Get("/", handlers.GetTasks)
	r.Post("/create", handlers.CreateTask)
	r.Put("/update", handlers.UpdateTask)
	r.Delete("/delete", handlers.DeleteTask)

	// Wrap the mux with CORS middleware
	return r
}
