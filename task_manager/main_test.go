package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"task_manager/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

func init() {
	r = router.SetupRouter()
}
func TestGetTasks(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task 1")
}

func TestGetTaskByID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task 1")
}

func TestCreateTask(t *testing.T) {
	task := map[string]string{
		"title":       "New Task",
		"description": "Details of the new task",
		"due_date":    "2024-09-01",
		"status":      "pending",
	}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Task")
}

func TestUpdateTask(t *testing.T) {
	task := map[string]string{
		"title":       "Updated Task",
		"description": "Updated details of the task",
		"due_date":    "2024-09-01",
		"status":      "completed",
	}
	taskJSON, _ := json.Marshal(task)

	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Task")
}

func TestDeleteTask(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
