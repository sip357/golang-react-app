// handlers_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"go-project/go-backend/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateTask tests the CreateTask handler
func TestCreateTask(t *testing.T) {
	// Sample input task
	task := models.Task{
		Title:   "New Task",
		Content: "This is a new task",
	}

	// Convert the task to JSON
	taskJSON, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Error marshalling task: %v", err)
	}

	// Create a new request with the task data
	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a new ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %v but got %v", http.StatusCreated, rr.Code)
	}

	// Check the response body
	var responseTask models.Task
	if err := json.NewDecoder(rr.Body).Decode(&responseTask); err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if responseTask.Title != task.Title {
		t.Errorf("Expected task title %v but got %v", task.Title, responseTask.Title)
	}
}
