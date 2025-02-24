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

func AuthCreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthCreate Task Called")

	//Ensure method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task models.AuthTask
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//Connect to Database in MongoDB
	client, nil := utils.ConnectToDB()
	coll := client.Database(os.Getenv("PROJECT_DB")).Collection(os.Getenv("TASK_COLLECTION"))

	//Insert user to database
	result, err := coll.InsertOne(context.TODO(), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Printf("Task created successfully _id: %v\n", result.InsertedID)
	json.NewEncoder(w).Encode("Task created successfully!")
}
