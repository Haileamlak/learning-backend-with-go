package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task_manager/models"
	"task_manager/router"
)

func TestCreateTask(t *testing.T) {
	r := router.SetupRouter()

	task := models.Task{
		Title:       "New Task",
		Description: "This is a new task",
		DueDate:     "2024-08-06",
		Status:      "pending",
	}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}

	if response["message"] != "Task created successfully" {
		t.Errorf("Expected message 'Task created successfully', got %s", response["message"])
	}
}
func TestGetAllTasks(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var tasks []models.Task
	if err := json.Unmarshal(w.Body.Bytes(), &tasks); err != nil {
		t.Errorf("Error unmarshaling response: %v", err)
	}
}

func TestGetTaskByID(t *testing.T) {
	r := router.SetupRouter()

	// First, create a new task to get by ID
	task := models.Task{
		Title:       "New Task",
		Description: "This is a new task",
		DueDate:     "2024-08-06",
		Status:      "pending",
	}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error unmarshaling response: %v", err)
	}

	taskID, ok := response["task"].(map[string]interface{})["InsertedID"].(string)
	if !ok {
		t.Fatalf("Expected task ID in response")
	}

	// Now, get the task by ID
	req, _ = http.NewRequest("GET", "/tasks/"+taskID, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var taskResponse map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &taskResponse); err != nil {
		t.Fatalf("Error unmarshaling response: %v", err)
	}

	if taskResponse["title"] != "New Task" {
		t.Errorf("Expected task title 'New Task', got %s", taskResponse["title"])
	}
}

func TestUpdateTask(t *testing.T) {
	r := router.SetupRouter()

	// First, create a new task to update
	task := models.Task{
		Title:       "New Task",
		Description: "This is a new task",
		DueDate:     "2024-08-06",
		Status:      "pending",
	}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error unmarshaling response: %v", err)
	}

	taskID, ok := response["task"].(map[string]interface{})["InsertedID"].(string)
	if !ok {
		t.Fatalf("Expected task ID in response")
	}

	// Now, update the task
	updatedTask := models.Task{
		Title:       "Updated Task",
		Description: "This is an updated task",
		DueDate:     "2024-08-06",
		Status:      "completed",
	}
	updatedTaskJSON, _ := json.Marshal(updatedTask)

	req, _ = http.NewRequest("PUT", "/tasks/"+taskID, bytes.NewBuffer(updatedTaskJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	json.Unmarshal(w.Body.Bytes(), &response)
	if response["message"] != "Task updated successfully" {
		t.Errorf("Expected message 'Task updated successfully', got %s", response["message"])
	}
}

func TestDeleteTask(t *testing.T) {
	r := router.SetupRouter()

	// First, create a new task to delete
	task := models.Task{
		Title:       "New Task",
		Description: "This is a new task",
		DueDate:     "2024-08-06",
		Status:      "pending",
	}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	taskID, ok := response["task"].(map[string]interface{})["InsertedID"].(string)
	
	if !ok {
		t.Fatalf("Expected task ID in response")
	}

	// Now, delete the task
	req, _ = http.NewRequest("DELETE", "/tasks/"+taskID, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	json.Unmarshal(w.Body.Bytes(), &response)
	if response["message"] != "Task deleted successfully" {
		t.Errorf("Expected message 'Task deleted successfully', got %s", response["message"])
	}
}
