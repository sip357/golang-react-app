package users

import (
	"context"
	"encoding/json"
	"fmt"
	"go-project/go-backend/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser allows user to sign up
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser handler called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Created = time.Now()

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	// Convert hashed password to a string and assign it back to the user object
	user.Password = string(hashedPassword)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(os.Getenv("USER_DB")).Collection(os.Getenv("USER_COLLECTION"))

	//Insert user to database
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("User created successfully _id: %v\n", result.InsertedID)
	json.NewEncoder(w).Encode("User created successfully!")
}
