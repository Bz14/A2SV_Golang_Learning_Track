package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

var taskManager = data.NewTaskManager()
var users = data.UserCollection()

/* GET /tasks all tasks */
func AllTaskHandler(ctx *gin.Context){
	role, user_id, err := userInfo(ctx)
	if err != nil{
		ctx.IndentedJSON(http.StatusUnauthorized, 
			gin.H{
				"message" : "Access denied",
		})
		return
	}
	tasks := taskManager.GetAllTasks(role, user_id)
	if tasks != nil{
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"tasks" : tasks,
		})
	}else{
		ctx.IndentedJSON(http.StatusNotFound, 
			gin.H{
				"message" : "Task not found",
		})
	}
}

/* GET /tasks:id task by ID */
func TaskByIdHandler(ctx *gin.Context){
	role, user_id, err := userInfo(ctx)
	if err != nil{
		ctx.IndentedJSON(http.StatusUnauthorized, 
			gin.H{
				"message" : "Access denied",
		})
		return
	}
	id := ctx.Param("id")
	task := taskManager.GetTaskById(id, user_id, role)
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
	role, user_id, err := userInfo(ctx)
	if err != nil{
		ctx.IndentedJSON(http.StatusUnauthorized, 
			gin.H{
				"message" : "Access denied",
		})
		return
	}
	id := ctx.Param("id")
	found := taskManager.DeleteTaskById(id, user_id, role)
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
	_, user_id, err := userInfo(ctx)
	if err != nil{
		ctx.IndentedJSON(http.StatusUnauthorized, 
			gin.H{
				"message" : "Access denied",
		})
		return
	}
	var task models.Task
	err = ctx.ShouldBind(&task)
	fmt.Println(err)
	if err != nil{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Task not created",
		})
		return
	}
	taskId := taskManager.CreateTask(task, user_id)
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
	role, user_id, err := userInfo(ctx)
	if err != nil{
		ctx.IndentedJSON(http.StatusUnauthorized, 
			gin.H{
				"message" : "Access denied",
		})
		return
	}
	result, err := taskManager.UpdateTask(id, user_id, role, task)
	if result {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			 "message" : "Task Updated",
		})
	}else if err != nil{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "Task not found",
		})
	}else{
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message" : "Task not found",
		})
	}

}

func UserRegisterHandler(c *gin.Context){
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "User not created",
		})
	}
	result, err := users.Register(user)
	if err == nil{
		c.JSON(http.StatusCreated, gin.H{
			"message" : "User registered successfully.",
			"userId" : result,
		})
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
	}
}

func UserLoginHandler(c *gin.Context){
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Invalid credentials.",
		})
	}
	token, err := users.Login(user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"message" : "User logged successfully.",
			"token" : token,
		})
	}
}

func userInfo(ctx *gin.Context)(string, string, error){
	role, exists := ctx.Get("role")
	if !exists{
		return "", "", errors.New("access denied")
	}
	if strings.ToUpper(role.(string)) == "ADMIN"{
		return "ADMIN", "", nil
	}
	strRole, ok := role.(string)
	if !ok {
		return "", "", errors.New("access denied")
	}
	userID, exists := ctx.Get("user_id")
	if !exists {
		return "", "", errors.New("user not logged in")
	}
	strId, ok := userID.(string)
	if !ok {
		return "", "", errors.New("user not logged in")
	}
	return strRole, strId, nil
}