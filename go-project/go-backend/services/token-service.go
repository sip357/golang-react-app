package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Replace with an environment variable in production!

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID string) (string, error) {
	// Create the claims (payload)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(), // Expires in 1 hour
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates a JWT token and returns the claims if valid
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// AuthError represents different types of authentication errors.
var (
	ErrNoToken       = errors.New("no token found")
	ErrInvalidToken  = errors.New("token is not valid")
	ErrTokenExpired  = errors.New("token has expired")
)

// AuthenticateToken validates a JWT token, checks expiration, and returns claims if valid.
func AuthenticateToken(token string) (map[string]interface{}, error) {
	// Validate the token
	claims, err := ValidateJWT(token)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// Check if the token has expired
	if exp, ok := claims["exp"].(float64); ok {
		expiryTime := int64(exp)
		if expiryTime < time.Now().Unix() {
			return nil, ErrTokenExpired
		}
	} else {
		return nil, ErrInvalidToken
	}

	return claims, nil
}