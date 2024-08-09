package controllers

import (
	"errors"
	"net/http"
	"strings"
	domain "task_manager/Domain"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userUseCase domain.UserUseCase
	taskUseCase domain.TaskUseCase
}

func NewController(userCase domain.UserUseCase, taskCase domain.TaskUseCase)*Controller{
	return &Controller{
		userUseCase: userCase,
		taskUseCase: taskCase,
	}
}

func (controller *Controller) GetAllTaskHandler(c *gin.Context){
	var result interface{}
	role, userId, err := userInfo(c)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	allTasks, err := controller.taskUseCase.GetAllTasks(role, userId)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	if len(allTasks) == 0{
		result = "Task not found"
	}else{
		result = allTasks
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks" : result,
	})
}

func (controller *Controller) TaskByIdHandler(c *gin.Context){
	id := c.Param("id")
	role, userId, err := userInfo(c)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	task, err := controller.taskUseCase.GetTaskById(id, userId, string(role))
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"task" : task,
	})
}

func (controller *Controller) DeleteTaskHandler(c *gin.Context){
	id := c.Param("id")
	role, userId, err := userInfo(c)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	deleted, err := controller.taskUseCase.DeleteTaskById(id, userId, role)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"message" : err.Error(),
		})
		return
	}
	if deleted{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : "Task deleted",
		})
	}else{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : "Task not found",
		})
	}
	
}

func (controller *Controller) CreateTaskHandler(c *gin.Context){
	var task domain.Task
	err := c.ShouldBind(&task)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	role, userId, err := userInfo(c)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	createdTask, err := controller.taskUseCase.CreateTask(task, userId, role)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message" : "Task created successfully.",
		"task_id" : createdTask,
	})
}

func (controller *Controller) UpdateTaskHandler(c *gin.Context){
	var task domain.Task
	err := c.ShouldBind(&task)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Invalid credentials",
		})
		return 
	}
	id := c.Param("id")
	role, userId, err := userInfo(c)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return 
	}
	updated, err := controller.taskUseCase.UpdateTask(id, userId, role, task)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})
		return
	}
	if !updated{
		c.JSON(http.StatusOK, gin.H{
			"message" : "Task not found", 
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Task updated successfully", 
	})
}

func (controller *Controller) UserRegisterHandler(c *gin.Context){
	var user domain.User
	err := c.ShouldBind(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Invalid credential", 
		})
		return
	}
	signedUser, err := controller.userUseCase.Register(user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(), 
		})
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Logged successfully",
			"user_id" : signedUser, 
		})
	}

}

func (controller *Controller) UserLoginHandler(c *gin.Context){
	var user domain.User
	err := c.ShouldBind(&user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Invalid credential", 
		})
		return
	}
	loggedUser, err := controller.userUseCase.Login(user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(), 
		})
	}else{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Logged successfully",
			"token" : loggedUser, 
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
