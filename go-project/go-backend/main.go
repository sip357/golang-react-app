package main

import (
	"context"
	"fmt"
	"go-project/go-backend/routes" // Import routes
	"go-project/go-backend/utils"  // Import utils
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	utils.Connect() // Connect to the database

	// Create a new HTTP server instance
	srv := &http.Server{
		Addr: ":443",
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	certFile := os.Getenv("CRT_PATH")
	keyFile := os.Getenv("KEY_PATH")

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Fatalf("Certificate file not found: %s", certFile)
	}
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		log.Fatalf("Key file not found: %s", keyFile)
	}

	go func() {
		fmt.Println("Starting HTTP server to redirect to HTTPS...")
		err := http.ListenAndServe(":80", http.RedirectHandler("https://localhost:443", http.StatusMovedPermanently))
		if err != nil {
			log.Fatal("Error starting HTTP redirect server: ", err)
		}
	}()

	http.Handle("/", routes.TaskRoutes())       // Handles all task routes
	http.Handle("/users/", routes.UserRoutes()) // Handles all user routes

	// Start the server
	fmt.Println("Server is running at https://localhost:443")
	go func() {
		if err := http.ListenAndServeTLS(srv.Addr, certFile, keyFile, nil); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Create a context with a timeout to allow graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown failed:", err)
	}
	fmt.Println("Server gracefully stopped")
}
