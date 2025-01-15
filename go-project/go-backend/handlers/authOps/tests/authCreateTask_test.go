// authCreateTask_test.go

package authtest

import (
	"bytes"
	"context"
	"encoding/json"
	auth "go-project/go-backend/handlers/authOps"
	"go-project/go-backend/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock database functions
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, doc interface{}) (interface{}, error) {
	args := m.Called(ctx, doc)
	return args.Get(0), args.Error(1)
}

func TestAuthCreateTask(t *testing.T) {
	// Mock environment variables
	if err := godotenv.Load(); err != nil {
		t.Fatalf("Error: %s", err)
	}

	//Error = .env file not found
	os.Setenv("PROJECT_DB", "test_db")
	os.Setenv("TASK_COLLECTION", "test_collection")

	// Mock database connection
	mockColl := new(MockCollection)
	mockColl.On("InsertOne", mock.Anything, mock.Anything).Return(primitive.NewObjectID(), nil)

	// Prepare input
	var username = "sip2345"
	var user_id *string = &username

	//Current time
	currentTime := time.Now()

	//Test Task variable
	task := models.AuthTask{
		Title:   "First User Task",
		Content: "This is the first authenticated task",
		UserID:  user_id,
		Created: &currentTime,
	}
	body, _ := json.Marshal(task)

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodPost, "/auth/task", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// Call the handler
	auth.AuthCreateTask(rec, req)

	// Validate response
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "Task created successfully")
	mockColl.AssertExpectations(t)
}
