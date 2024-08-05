package controller

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

var taskManager = data.NewTaskManager()

/* GET /tasks all tasks */
func AllTaskHandler(ctx *gin.Context){
	tasks := taskManager.GetAllTasks()
	if tasks != nil{
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"tasks" : tasks,
		})
	}else{
		ctx.IndentedJSON(http.StatusNotFound, 
			gin.H{
				"message" : "Error fetching on tasks",
		})
	}
}

/* GET /tasks:id task by ID */
func TaskByIdHandler(ctx *gin.Context){
	id := ctx.Param("id")
	task := taskManager.GetTaskById(id)
	if task != nil{
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
	found := taskManager.DeleteTaskById(id)
	if found != 0{
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message" : "Task deleted",
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
	taskId := taskManager.CreateTask(task)
	if taskId != nil{
		ctx.IndentedJSON(http.StatusCreated, 
			gin.H{
				"message" : "Task Created",
				"task" : taskId,
			})
	}else{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Task not created",
		})
	}
	
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
	result := taskManager.UpdateTask(id, task)
	if result {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			 "message" : "Task Updated",
		})
	}else{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "Task not found",
		})
	}
}
