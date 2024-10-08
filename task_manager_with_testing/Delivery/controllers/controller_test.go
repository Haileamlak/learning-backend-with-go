package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	domain "task-manager/Domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) GetTask(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}	

func (m *MockTaskUsecase) UpdateTask(id string, task domain.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}


func (m *MockTaskUsecase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(username, password string) error {
	args := m.Called(username, password)
	return args.Error(0)
}

func (m *MockUserUsecase) Login(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) PromoteUser(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

type ApiControllerTestSuite struct {
	suite.Suite
	taskUsecase *MockTaskUsecase
	userUsecase *MockUserUsecase
	controller  ApiController
	router      *gin.Engine
}

func (suite *ApiControllerTestSuite) SetupTest() {
	suite.taskUsecase = new(MockTaskUsecase)
	suite.userUsecase = new(MockUserUsecase)
	suite.controller = NewApiController(suite.taskUsecase, suite.userUsecase)
	suite.router = gin.Default()

	// Register routes
	suite.router.POST("/tasks", suite.controller.CreateTask)
	suite.router.GET("/tasks/:id", suite.controller.GetTask)
	suite.router.GET("/tasks", suite.controller.GetTasks)
	suite.router.PUT("/tasks/:id", suite.controller.UpdateTask)
	suite.router.DELETE("/tasks/:id", suite.controller.DeleteTask)
	suite.router.POST("/register", suite.controller.Register)
	suite.router.POST("/login", suite.controller.Login)
	suite.router.POST("/promote", suite.controller.PromoteUser)
}

func TestApiControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ApiControllerTestSuite))
}

func (suite *ApiControllerTestSuite) TestCreateTask_Success() {
	dueDate, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	task := domain.Task{Title: "Test Task", DueDate: dueDate, Status: "pending"}
	suite.taskUsecase.On("CreateTask", task).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(`{"title": "Test Task", "due_date": "2021-01-01T00:00:00Z", "status": "pending"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Task created successfully")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestCreateTask_BadRequest() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(`{"title": ""}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag")
	suite.taskUsecase.AssertNotCalled(suite.T(), "CreateTask", mock.Anything)
}

func (suite *ApiControllerTestSuite) TestCreateTask_Error() {
	dueDate, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	task := domain.Task{Title: "Test Task", DueDate: dueDate, Status: "pending"}
	suite.taskUsecase.On("CreateTask", task).Return(&domain.InternalServerError{Message: "Internal server error"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(`{"title": "Test Task", "due_date": "2021-01-01T00:00:00Z", "status": "pending"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Internal server error")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestGetTask_Success() {
	task := domain.Task{Title: "Test Task", DueDate: time.Now().Add(24 * time.Hour), Status: "pending"}
	suite.taskUsecase.On("GetTask", "1").Return(task, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Test Task")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestGetTask_NotFound() {
	suite.taskUsecase.On("GetTask", "1").Return(domain.Task{}, &domain.NotFoundError{Message: "Task not found"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	suite.router.ServeHTTP(w, req)
	
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Task not found")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestGetTask_Error() {
	suite.taskUsecase.On("GetTask", "1").Return(domain.Task{}, &domain.InternalServerError{Message: "Internal server error"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Internal server error")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestGetTasks_Success() {
	tasks := []domain.Task{
		{Title: "Test Task 1", DueDate: time.Now().Add(24 * time.Hour), Status: "pending"},
		{Title: "Test Task 2", DueDate: time.Now().Add(48 * time.Hour), Status: "completed"},
	}
	suite.taskUsecase.On("GetTasks").Return(tasks, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Test Task 1")
	assert.Contains(suite.T(), w.Body.String(), "Test Task 2")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestGetTasks_Error() {
	suite.taskUsecase.On("GetTasks").Return([]domain.Task{}, &domain.InternalServerError{Message: "Internal server error"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Internal server error")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestUpdateTask_Success() {
	dueDate, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	task := domain.Task{Title: "Test Task", DueDate: dueDate, Status: "pending"}
	suite.taskUsecase.On("UpdateTask", "1", task).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/1", strings.NewReader(`{"title": "Test Task", "due_date": "2021-01-01T00:00:00Z", "status": "pending"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Task updated successfully")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestUpdateTask_BadRequest() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/1", strings.NewReader(`{"title": ""}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag")
	suite.taskUsecase.AssertNotCalled(suite.T(), "UpdateTask", mock.Anything, mock.Anything)
}

func (suite *ApiControllerTestSuite) TestUpdateTask_Error() {
	dueDate, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	task := domain.Task{Title: "Test Task", DueDate: dueDate, Status: "pending"}
	suite.taskUsecase.On("UpdateTask", "1", task).Return(&domain.InternalServerError{Message: "Internal server error"})
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/1", strings.NewReader(`{"title": "Test Task", "due_date": "2021-01-01T00:00:00Z", "status": "pending"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Internal server error")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestDeleteTask_Success() {
	suite.taskUsecase.On("DeleteTask", "1").Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Task deleted successfully")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestDeleteTask_Error() {
	suite.taskUsecase.On("DeleteTask", "1").Return(&domain.InternalServerError{Message: "Internal server error"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Internal server error")
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestRegister_Success() {
	suite.userUsecase.On("Register", "testuser", "password").Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{"username": "testuser", "password": "password"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "User registered successfully")
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestRegister_BadRequest() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{"username": ""}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag")
	suite.userUsecase.AssertNotCalled(suite.T(), "Register", mock.Anything, mock.Anything)
}

func (suite *ApiControllerTestSuite) TestRegister_Error() {
	suite.userUsecase.On("Register", "testuser", "password").Return(&domain.InternalServerError{Message: "Internal server error"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{"username": "testuser", "password": "password"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Internal server error")
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestLogin_Success() {
	suite.userUsecase.On("Login", "testuser", "password").Return("token", nil)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"username": "testuser", "password": "password"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "token")
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestLogin_BadRequest() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"username": ""}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag")
	suite.userUsecase.AssertNotCalled(suite.T(), "Login", mock.Anything, mock.Anything)
}

func (suite *ApiControllerTestSuite) TestLogin_Error() {
	suite.userUsecase.On("Login", "testuser", "password").Return("", &domain.InternalServerError{Message: "Internal server error"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"username": "testuser", "password": "password"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Internal server error")
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestPromoteUser_Success() {
	suite.userUsecase.On("PromoteUser", "testuser").Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/promote", strings.NewReader(`{"username": "testuser"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "User promoted successfully")
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *ApiControllerTestSuite) TestPromoteUser_BadRequest() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/promote", strings.NewReader(`{"username": ""}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Key: 'Username' Error:Field validation for 'Username' failed on the 'required' tag")
	suite.userUsecase.AssertNotCalled(suite.T(), "PromoteUser", mock.Anything)
}

func (suite *ApiControllerTestSuite) TestPromoteUser_Error() {
	suite.userUsecase.On("PromoteUser", "testuser").Return(&domain.InternalServerError{Message: "Internal server error"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/promote", strings.NewReader(`{"username": "testuser"}`))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "Internal server error")
	suite.userUsecase.AssertExpectations(suite.T())
}
