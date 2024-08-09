package main

import (
	"task_manager/Delivery/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	routers.Route(route)
	route.Run()
}