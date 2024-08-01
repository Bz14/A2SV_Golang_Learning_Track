package controller

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func AllTaskHandler(ctx *gin.Context){
	allData := data.GetAllTasks()
	ctx.IndentedJSON(http.StatusOK, allData)
}

func TaskByIdHandler(ctx *gin.Context){
	id := ctx.Param("id")
	found, task := data.GetTaskById(id)
	if found{
		ctx.IndentedJSON(http.StatusOK, task)
	}else{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "Task Not Found",
		})
	}
}

func DeleteTaskHandler(ctx *gin.Context){
	id := ctx.Param("id")
	found := data.DeleteTaskById(id)
	if found{
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message" : "Task Deleted",
		})
	}else{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "Task not found",
		})
	}
}


func CreateTaskHandler(ctx *gin.Context){
	var task models.Task
	err := ctx.ShouldBind(&task)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Can not create a task",
		})
		return
	}
	data.CreateTask(task)
	ctx.IndentedJSON(http.StatusCreated, 
		gin.H{
			"message" : "Task Created",
			"task" : task,
		})
}


func UpdateTaskHandler(ctx *gin.Context){
	var task models.Task
	id := ctx.Param("id")
	err := ctx.ShouldBind(&task)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "Can not updated file.",
		})
	}
	if data.UpdateTask(id, task){
		ctx.IndentedJSON(http.StatusOK, gin.H{
			 "message" : "Task Updated.",
			 "task" : task,
		})
	}else{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "Task not found",
		})
	}
}