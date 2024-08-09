package routers

import (
	controllers "task_manager/Delivery/Controllers"
	database "task_manager/Delivery/Database"
	domain "task_manager/Domain"
	infrastructure "task_manager/Infrastructure"
	repository "task_manager/Repositories"
	usecases "task_manager/Usecases"

	"github.com/gin-gonic/gin"
)

func Setup()(*controllers.Controller, *infrastructure.Auth){
	db := database.NewDatabase()
	userCollection := db.CreateDb(domain.CollectionUser)
	taskCollection := db.CreateDb(domain.CollectionTask)

	password := infrastructure.NewPasswordHash()
	jwtToken := infrastructure.NewJWtService()

	userRepository := repository.NewUserRepository(userCollection, password, jwtToken)
	taskRepository := repository.NewTaskRepository(taskCollection)

	userUseCase := usecases.NewUserUseCase(userRepository)
	taskUseCase := usecases.NewTaskUseCase(taskRepository)

	controller := controllers.NewController(userUseCase, taskUseCase)

	middleware := infrastructure.NewAuthMiddleware(jwtToken)
	return controller, middleware
}


func Route(route *gin.Engine){
	controller, middleware := Setup()
	route.GET("/tasks", middleware.AuthenticationMiddleware(), controller.GetAllTaskHandler)
	route.GET("/tasks/:id", middleware.AuthenticationMiddleware(), controller.TaskByIdHandler)
	route.DELETE("/tasks/:id",middleware.AuthenticationMiddleware(), controller.DeleteTaskHandler)
	route.POST("/tasks", middleware.AuthenticationMiddleware(), controller.CreateTaskHandler)
	route.PUT("/tasks/:id", middleware.AuthenticationMiddleware(), controller.UpdateTaskHandler)

	route.POST("/register", controller.UserRegisterHandler)
	route.POST("/login", controller.UserLoginHandler)
}