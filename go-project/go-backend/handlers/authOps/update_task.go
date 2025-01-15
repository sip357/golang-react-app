// Update task function can only be used by an authenticated user
package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go-project/go-backend/models" //Task Models
	"go-project/go-backend/utils"  //Database connection
)

func AuthUpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthUpdate Task Called")

	//Ensure method is POST
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedTask models.AuthTask
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//Connect to Database in MongoDB
	client, nil := utils.ConnectToDB()
	coll := client.Database(os.Getenv("PROJECT_DB")).Collection(os.Getenv("TASK_COLLECTION"))

	//Update task in database
	result, err := coll.UpdateByID(context.TODO(), updatedTask.ID, updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("Task Updated Successfully: %v\n", result.UpsertedID)
	json.NewEncoder(w).Encode("Task Updated Successfully!")
}
