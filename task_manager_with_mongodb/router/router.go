package router

import (
    "task_manager/data"
    "task_manager/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

	taskService := data.NewTaskService()
	controller := controllers.NewTaskController(taskService)

    r.POST("/tasks", controller.CreateTask)
    r.GET("/tasks", controller.GetTasks)
    r.GET("/tasks/:id", controller.GetTaskByID)
    r.PUT("/tasks/:id", controller.UpdateTask)
    r.DELETE("/tasks/:id", controller.DeleteTask)

    return r
}
