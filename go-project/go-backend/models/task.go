package models

import "time"

type Task struct {
	ID      int        `json:"id"`
	Title   string     `json:"title"`
	Content string     `json:"content"`
	Created *time.Time `json:"created_at,omitempty"`
}

type AuthTask struct {
	ID      int        `json:"id,omitempty" bson:"_id,omitempty"`
	Title   string     `json:"title" bson:"title"`
	Content string     `json:"content" bson:"content"`
	UserID  *string    `json:"user_id,omitempty" bson:"user_id,omitempty"` // optional field that stores the username of the client
	Created *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
