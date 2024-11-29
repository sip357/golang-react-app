package main

import (
	"fmt"
	"go-project/go-backend/routes" // Import routes
	"net/http"
)

func main() {
	// Register task routes
	http.Handle("/", routes.TaskRoutes()) // Handles all "/tasks/*" routes

	// Register user routes
	http.Handle("/users/", routes.UserRoutes()) // Handles all "/users/*" routes

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
