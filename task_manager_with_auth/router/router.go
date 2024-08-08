package router

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	taskService := data.NewTaskService()
	userService := data.NewUserService()
	controller := controllers.NewApiController(taskService, userService)

	r.POST("/register", controller.RegisterUser)
	r.POST("/login", controller.LoginUser)

	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())

	adminMiddleware := middleware.AdminMiddleware()
	
	authorized.POST("/tasks", adminMiddleware, controller.CreateTask)
	authorized.PUT("/tasks/:id", adminMiddleware, controller.UpdateTask)
	authorized.DELETE("/tasks/:id", adminMiddleware, controller.DeleteTask)
	authorized.GET("/tasks", controller.GetTasks)
	authorized.GET("/tasks/:id", controller.GetTaskByID)
	authorized.POST("/promote/:id", adminMiddleware, controller.PromoteUser)

	return r
}
