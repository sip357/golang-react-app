package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-project/go-backend/models"            //Task Models
	services "go-project/go-backend/services" //Task Services
)

// Create Task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateTask handler called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task = services.AddTaskService(task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
