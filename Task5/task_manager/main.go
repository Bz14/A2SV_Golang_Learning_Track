package main

import (
	"task_manager/router"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	router.Router(route)
	route.Run()
}