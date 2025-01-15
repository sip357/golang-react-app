package users

import (
	"encoding/json"
	"fmt"
	"go-project/go-backend/models"
	services "go-project/go-backend/task-service"
	"go-project/go-backend/utils"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser allows user to register
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser handler called")

	// Ensure method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActive = true
	emptyString := ""
	user.ResetToken = &emptyString

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	// Convert hashed password to a string and assign it back to the user object
	user.Password = string(hashedPassword)

	//Generate API Key
	apiKey := services.GenerateAPIKey()
	user.APIKey = &apiKey

	//Load environment variables
	if err := godotenv.Load(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error loading .env file") //Delete this line in production
		log.Println("No .env file found")
		return
	}

	//Connect to Database in MySQL
	db, err := utils.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error connecting to database: %v\n", err)
		json.NewEncoder(w).Encode("Error connecting to database")
		return
	}

	//Close the database connection
	defer db.Close()

	//Check if username already exists
	exists, err := services.UsernameExists(db, user.Username)
	if exists {
		w.WriteHeader(http.StatusConflict)
		fmt.Printf("Username already exists: %v\n", user.Username)
		json.NewEncoder(w).Encode("Username already exists")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error checking if username exists: %v\n", err)
		json.NewEncoder(w).Encode("Error checking if username exists")
		return
	}

	//Check if email already exists
	exists, err = services.EmailExists(db, user.Email)
	if exists {
		w.WriteHeader(http.StatusConflict)
		fmt.Printf("Email already exists: %v\n", user.Email)
		json.NewEncoder(w).Encode("Email already exists")
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error checking if email exists: %v\n", err)
		json.NewEncoder(w).Encode("Error checking if email exists")
		return
	}

	//Insert user into the database
	query := `INSERT INTO users (username, email, password_hash, created_at, updated_at, is_active, reset_token, api_key) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	//Execute the query and check for errors
	result, err := db.Exec(query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.IsActive, user.ResetToken, user.APIKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Query error: %v\n", err)
		return
	}

	//Return a success status
	w.WriteHeader(http.StatusCreated)

	insertedID, err := result.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error getting last inserted ID: %v\n", err)
		return
	}

	//Print the result and user
	fmt.Printf("User created successfully _id: %v\n", insertedID)

	//Return a success message
	json.NewEncoder(w).Encode("User created successfully!")
}
