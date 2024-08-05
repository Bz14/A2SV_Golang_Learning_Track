package router

import (
	controller "task_manager/controllers"

	"github.com/gin-gonic/gin"
)

/* Routing Http Requests */
func Router(route *gin.Engine) {
	route.GET("/tasks", controller.AllTaskHandler)
	route.GET("/tasks/:id", controller.TaskByIdHandler)
	route.DELETE("/tasks/:id", controller.DeleteTaskHandler)
	route.POST("/tasks", controller.CreateTaskHandler)
	route.PUT("/tasks/:id", controller.UpdateTaskHandler)
}