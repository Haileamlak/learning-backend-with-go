package router

import (
    "task_manager/controllers"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/tasks", controllers.GetTasks)
    r.GET("/tasks/:id", controllers.GetTaskByID)
    r.PUT("/tasks/:id", controllers.UpdateTask)
    r.POST("/tasks", controllers.CreateTask)
    r.DELETE("/tasks/:id", controllers.DeleteTask)

    return r
}