package models

import "time"

// User struct maps to the 'users' table in MySQL
type User struct {
	ID         int       `json:"id"`                    // Maps to id (INT AUTO_INCREMENT)
	Username   string    `json:"username"`              // Maps to username (VARCHAR)
	Email      string    `json:"email"`                 // Maps to email (VARCHAR)
	Password   string    `json:"password_hash"`         // Maps to password_hash (VARCHAR)
	Role       string    `json:"role"`                  // Maps to role (ENUM), user or admin, user is default
	IsActive   bool      `json:"is_active"`             // Maps to is_active (BOOLEAN)
	ResetToken *string   `json:"reset_token,omitempty"` // Maps to reset_token (VARCHAR, nullable)
	APIKey     *string   `json:"api_key,omitempty"`     // Maps to api_key (VARCHAR, nullable)
	CreatedAt  time.Time `json:"created_at"`            // Maps to created_at (TIMESTAMP)
	UpdatedAt  time.Time `json:"updated_at"`            // Maps to updated_at (TIMESTAMP)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
