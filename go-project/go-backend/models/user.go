package models

import "time"

type User struct {
	ID       string    `json:"id,omitempty" bson:"_id,omitempty"`
	Username string    `json:"username" bson:"username"`
	Password string    `json:"password" bson:"password"`
	Email    string    `json:"email" bson:"email"`
	Created  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
