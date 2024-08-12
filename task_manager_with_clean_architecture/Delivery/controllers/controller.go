package controllers

import (
	"net/http"

	domain "task-manager/Domain"
	usecases "task-manager/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ApiController interface
type ApiController interface {
	CreateTask(c *gin.Context)
	GetTask(c *gin.Context)
	GetTasks(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	PromoteUser(c *gin.Context)
}

// apiController struct
type apiController struct {
	taskUsecase usecases.TaskUsecase
	userUsecase usecases.UserUsecase
}

// NewApiController creates a new api controller
func NewApiController(taskUsecase usecases.TaskUsecase, userUsecase usecases.UserUsecase) ApiController {
	return &apiController{taskUsecase, userUsecase}
}

// CreateTask creates a new task
func (c *apiController) CreateTask(ctx *gin.Context) {
	task := domain.Task{}
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.taskUsecase.CreateTask(task)
	if err != nil {
		switch err.(type) {
		case *domain.BadRequestError:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case *domain.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

// GetTask retrieves a task by ID
func (c *apiController) GetTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := c.taskUsecase.GetTask(id)
	if err != nil {
		switch err.(type) {
		case *domain.NotFoundError:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case *domain.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// GetTasks retrieves all tasks
func (c *apiController) GetTasks(ctx *gin.Context) {
	tasks, err := c.taskUsecase.GetTasks()
	if err != nil {
		switch err.(type) {
		case *domain.NotFoundError:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case *domain.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// UpdateTask updates a task
func (c *apiController) UpdateTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := domain.Task{}
	err = ctx.BindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.taskUsecase.UpdateTask(id, task)
	if err != nil {
		switch err.(type) {
		case *domain.BadRequestError:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case *domain.NotFoundError:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case *domain.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// DeleteTask deletes a task
func (c *apiController) DeleteTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.taskUsecase.DeleteTask(id)
	if err != nil {
		switch err.(type) {
		case *domain.NotFoundError:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case *domain.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// Register registers a new user
func (c *apiController) Register(ctx *gin.Context) {
	var registerInfo domain.User

	err := ctx.BindJSON(&registerInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.userUsecase.Register(registerInfo.Username, registerInfo.Password)
	if err != nil {
		switch err.(type) {
		case *domain.BadRequestError:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case *domain.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login logs in a user
func (c *apiController) Login(ctx *gin.Context) {
	var loginInfo domain.User

	err := ctx.BindJSON(&loginInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userUsecase.Login(loginInfo.Username, loginInfo.Password)
	if err != nil {
		switch err.(type) {
		case *domain.BadRequestError:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case *domain.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "token": token})
}

// PromoteUser promotes a user to admin
func (c *apiController) PromoteUser(ctx *gin.Context) {
	var userInfo struct {
		Username string `json:"username"`
	}
	err := ctx.BindJSON(&userInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.userUsecase.PromoteUser(userInfo.Username)
	if err != nil {
		switch err.(type) {
		case *domain.NotFoundError:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case *domain.InternalServerError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}