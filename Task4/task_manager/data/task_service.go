package data

import (
	"strconv"
	"task_manager/models"
	"time"
)

type TaskManager struct{
	listOfTasks map[string]models.Task
}

func NewTaskManager() *TaskManager{
	return &TaskManager{
		listOfTasks: map[string]models.Task{
			"1" : {ID: "1", Title: "Taking a note", Description: "Take Notes on the meeting", DueDate: time.Now(), Status: "Done"},
			"2" : {ID: "2", Title: "Sleeping", Description: "Sleep for two hours", DueDate: time.Now(), Status: "In Progress"},
		},
	}
}


var taskId = 3

/* List of all available tasks*/
func (t *TaskManager)GetAllTasks()[]models.Task{
	var tasks []models.Task
	for _, task := range t.listOfTasks{
		tasks = append(tasks, task)
	}
	return tasks
}

/* Get task by ID*/
func (t *TaskManager)GetTaskById(id string)(bool, models.Task){
	task, exists := t.listOfTasks[id]
	if !exists{
		return false, models.Task{}
	}
	return true, task
}

/* Delete task with a given ID */
func (t *TaskManager) DeleteTaskById(id string)bool{
	_, exists := t.listOfTasks[id]
	if !exists{
		return false
	}
	delete(t.listOfTasks, id)
	return true
}

/* Create a new task*/
func (t *TaskManager) CreateTask(newTask models.Task)models.Task{
	newTask.ID = strconv.Itoa(taskId)
	t.listOfTasks[newTask.ID] = newTask
	taskId++
	return newTask
}

/* Updating a task with a given ID*/
func  (t *TaskManager) UpdateTask(id string, newTask models.Task)bool{
	task, exists := t.listOfTasks[id]
	if !exists{
		return false
	}
	if newTask.Title != ""{
		task.Title =  newTask.Title
	}
	if newTask.Description != ""{
		task.Description =  newTask.Description
	}
	if newTask.Status != ""{
		task.Status =  newTask.Status
	}
	if !newTask.DueDate.IsZero(){
		task.DueDate =  newTask.DueDate
	}
	t.listOfTasks[id] = task
	return true
}