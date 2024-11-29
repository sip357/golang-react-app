package users

import (
	"context"
	"encoding/json"
	"fmt"
	"go-project/go-backend/models"
	services "go-project/go-backend/task-service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler authenticates a user and sets a JWT in an HTTP-only cookie
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginHandler called")

	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the login request
	var user models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Set up MongoDB connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.TODO())

	coll := client.Database("users").Collection("existing_users")

	// Find the user by username
	filter := bson.D{{Key: "username", Value: user.Username}}
	var existingUser models.User
	err = coll.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		return
	}

	// Generate the JWT
	token, err := services.GenerateJWT(existingUser.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Set the token as an HTTP-only cookie
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true, // Enable HTTPS in production
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
	}
	http.SetCookie(w, cookie)

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
