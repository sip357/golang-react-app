package auth

import (
	"fmt"
	services "go-project/go-backend/task-service"
	"net/http"
	"time"
)

// Ensures that the user is authenticated
func ProtectedRoute(w http.ResponseWriter, r *http.Request) {
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

	w.Write([]byte("Welcome your User ID is:" + claims["user_id"].(string)))
}
