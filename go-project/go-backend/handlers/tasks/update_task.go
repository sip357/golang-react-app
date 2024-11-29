package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-project/go-backend/models"                //Task Models
	services "go-project/go-backend/task-service" //Task Services
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateTask handler called")
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedTask, err := services.UpdateTaskService(updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// If no error, return the updated task as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}
