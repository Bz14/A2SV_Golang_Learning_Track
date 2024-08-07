package router

import (
	controller "task_manager/controllers"
	middleware "task_manager/middleware"

	"github.com/gin-gonic/gin"
)

/* Routing Http Requests */
func Router(route *gin.Engine) {
	route.GET("/tasks", middleware.AuthenticationMiddleware(), controller.AllTaskHandler)
	route.GET("/tasks/:id", middleware.AuthenticationMiddleware(), controller.TaskByIdHandler)
	route.DELETE("/tasks/:id",middleware.AuthenticationMiddleware(), controller.DeleteTaskHandler)
	route.POST("/tasks", middleware.AuthenticationMiddleware(), controller.CreateTaskHandler)
	route.PUT("/tasks/:id", middleware.AuthenticationMiddleware(), controller.UpdateTaskHandler)

	route.POST("/register", controller.UserRegisterHandler)
	route.POST("/login", controller.UserLoginHandler)
}