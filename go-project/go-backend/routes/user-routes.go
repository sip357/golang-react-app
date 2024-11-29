package routes

import (
	"go-project/go-backend/handlers/users"
	"net/http"

	"github.com/rs/cors" // CORS middleware package
)

// TaskRoutes with CORS middleware
func UserRoutes() http.Handler {
	mux := http.NewServeMux()

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Allow frontend from localhost:3000
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Routes associated with handlers

	mux.HandleFunc("/users/signUp", users.CreateUser)
	mux.HandleFunc("/users/login", users.LoginHandler)

	// Wrap the mux with CORS middleware
	return c.Handler(mux)
}
