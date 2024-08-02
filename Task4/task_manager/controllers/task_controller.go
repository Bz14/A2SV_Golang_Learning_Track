package controller

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

/* GET /tasks all tasks */
func AllTaskHandler(ctx *gin.Context){
	tasks := data.GetAllTasks()
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"tasks" : tasks,
	})
}

/* GET /tasks:id task by ID */
func TaskByIdHandler(ctx *gin.Context){
	id := ctx.Param("id")
	found, task := data.GetTaskById(id)
	if found{
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"task" : task,
		})
	}else{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "Task Not Found",
		})
	}
}

/* Delete tasks/:id delete task by ID*/
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

/* Post tasks/ create new task */
func CreateTaskHandler(ctx *gin.Context){
	var task models.Task
	err := ctx.ShouldBind(&task)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Task not created",
		})
		return
	}
	newTask := data.CreateTask(task)
	ctx.IndentedJSON(http.StatusCreated, 
		gin.H{
			"message" : "Task Created",
			"task" : newTask,
		})
}

/* Put tasks/:id update a task */
func UpdateTaskHandler(ctx *gin.Context){
	var task models.Task
	id := ctx.Param("id")
	err := ctx.ShouldBind(&task)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "Task not updated.",
		})
	}
	if data.UpdateTask(id, task){
		ctx.IndentedJSON(http.StatusOK, gin.H{
			 "message" : "Task Updated",
		})
	}else{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "Task not found",
		})
	}
}