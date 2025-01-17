package services

import (
	"fmt"
	"sync"

	"go-project/go-backend/models" //Task Model
)

var (
	tasks  = []models.Task{} //Array of Tasks
	nextID = 1
	mu     sync.Mutex
)

func AddTaskService(task models.Task) models.Task {
	mu.Lock()
	defer mu.Unlock() //Release resource at the end of function

	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	return task
}

func GetTasksService() []models.Task {
	mu.Lock()
	defer mu.Unlock()

	return tasks
}

func UpdateTaskService(updatedTask models.Task) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			return updatedTask, nil
		}
	}

	return models.Task{}, fmt.Errorf("task with ID %d not found", updatedTask.ID)
}

func DeleteTaskService(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}
