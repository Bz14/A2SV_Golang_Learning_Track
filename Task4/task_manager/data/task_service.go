package data

import (
	"strconv"
	"task_manager/models"
	"time"
)

var listOfTasks = []models.Task{
	{ID: "1", Title: "Taking a note", Description: "Take Notes on the meeting", DueDate: time.Now(), Status: "Done"},
	{ID: "2", Title: "Sleeping", Description: "Sleep for two hours", DueDate: time.Now(), Status: "In Progress"},
}

var taskId = 3

/* List of all available tasks*/
func GetAllTasks()[]models.Task{
	return listOfTasks
}

/* Get task by ID*/
func GetTaskById(id string)(bool, models.Task){
	for _, task := range listOfTasks{
		if task.ID == id{
			return true, task
		}
	}
	return false, models.Task{}
}

/* Delete task with a given ID */
func DeleteTaskById(id string)bool{
	var newListTask []models.Task
	found := false
	for _, task := range listOfTasks{
		if task.ID == id{
			found = true
			continue
		}
		newListTask = append(newListTask, task)
	}
	listOfTasks = newListTask
	return found
}

/* Create a new task*/
func CreateTask(newTask models.Task)models.Task{
	newTask.ID = strconv.Itoa(taskId)
	taskId++
	listOfTasks = append(listOfTasks, newTask)
	return newTask
}

/* Updating a task with a given ID*/
func UpdateTask(id string, task models.Task)bool{
	for i := 0; i < len(listOfTasks); i++{
		if listOfTasks[i].ID == id{
			if task.Title != ""{
				listOfTasks[i].Title = task.Title
			}
			if task.Description != ""{
				listOfTasks[i].Description = task.Description
			}
			if task.Status != ""{
				listOfTasks[i].Status = task.Status
			}
			return true
		}
	}
	return false
}