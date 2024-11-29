package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	//Task Models
	services "go-project/go-backend/task-service" //Task Services
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteTask handler called")
	// Get the 'id' parameter from the query string
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Call the DeleteTaskService
	err = services.DeleteTaskService(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// If successful, return No Content status
	w.WriteHeader(http.StatusNoContent)
}
