package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-project/go-backend/models"
	services "go-project/go-backend/services"
	"go-project/go-backend/utils"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
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

	//Set up MySQL connection
	db, err := utils.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}

	// Close the database connection when the function completes
	defer db.Close()

	//Find user by username
	query := "SELECT id, username, email, password_hash, is_active FROM users WHERE username = ?"
	row := db.QueryRow(query, user.Username)
	log.Println("Query Executed. Username:", user.Username) //Delete in production

	if row.Err() != nil {
		http.Error(w, "Couldn't select row", http.StatusUnauthorized)
		return
	}

	// Decode the user from the database
	var existingUser models.User
	err = row.Scan(&existingUser.ID, &existingUser.Username, &existingUser.Email, &existingUser.Password, &existingUser.IsActive)
	log.Println("User decoded", existingUser) //Delete in production
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusUnauthorized)
		} else {
			http.Error(w, "Couldn't scan row", http.StatusUnauthorized)
		}
		return
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		log.Printf("Password comparison failed: %v", err) // Log the error for debugging
		http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
		return
	}

	// Generate the JWT
	token, err := services.GenerateJWT(fmt.Sprintf("%d", existingUser.ID))
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
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)

	//Set is_active to true
	query = "UPDATE users SET is_active = ? WHERE id = ?"
	_, err = db.Exec(query, true, existingUser.ID)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
