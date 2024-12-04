package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go-project/go-backend/models"                //Task Models
	services "go-project/go-backend/task-service" //Task Services
	"go-project/go-backend/utils"                 //Database connection
)

func AuthCreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected Route Handler called")
	//get the cookie
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Error(w, "Unauthorized: no token found", http.StatusUnauthorized)
		return
	}

	//Get the token and verify it
	claims, err := services.ValidateJWT(cookie.Value)
	if err != nil {
		http.Error(w, "Token not valid", http.StatusUnauthorized)
		return
	}

	// Check if token has expired
	if exp, ok := claims["exp"].(float64); ok {
		// Convert expiration time to Unix timestamp
		expiryTime := int64(exp)

		if expiryTime < time.Now().Unix() {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "Token invalid", http.StatusUnauthorized)
		return
	}

	//Ensure method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	client, nil := utils.ConnectToDB()

	coll := client.Database(os.Getenv("PROJECT_DB")).Collection(os.Getenv("TASK_COLLECTION"))

	//Insert user to database
	result, err := coll.InsertOne(context.TODO(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("User created successfully _id: %v\n", result.InsertedID)
	json.NewEncoder(w).Encode("User created successfully!")
}
