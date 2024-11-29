package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	//Task Models
	services "go-project/go-backend/task-service" //Task Services
)

// Read All Tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetTasks handler called")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tasks := services.GetTasksService()
	json.NewEncoder(w).Encode(tasks)
}
