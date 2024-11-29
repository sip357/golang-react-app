package routes

import (
	handlers "go-project/go-backend/handlers/tasks" // import handlers

	"net/http"

	"github.com/rs/cors" // CORS middleware package
)

// TaskRoutes with CORS middleware
func TaskRoutes() http.Handler {
	mux := http.NewServeMux()

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Allow frontend from localhost:3000
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Routes associated with handlers
	mux.HandleFunc("/", handlers.GetTasks)
	mux.HandleFunc("/create", handlers.CreateTask)
	mux.HandleFunc("/update", handlers.UpdateTask)
	mux.HandleFunc("/delete", handlers.DeleteTask)

	// Wrap the mux with CORS middleware
	return c.Handler(mux)
}
