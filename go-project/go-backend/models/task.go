package models

import "time"

type Task struct {
	ID      int        `json:"id"`
	Title   string     `json:"title"`
	Content string     `json:"content"`
	UserID  *string    `json:"user_id,omitempty"` // optional field
	Created *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
