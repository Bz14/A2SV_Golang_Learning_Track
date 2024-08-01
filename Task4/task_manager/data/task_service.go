package data

import (
	"fmt"
	"task_manager/models"
)

var listOfTasks = []models.Task{
	{ID: "1", Title: "Taking a note", Description: "Take Notes on the meeting", DueDate: "Today", Status: "Done"},
	{ID: "2", Title: "Sleeping", Description: "Sleep for two hours", DueDate: "Today", Status: "In Progress"},
}

func GetAllTasks()[]models.Task{
	return listOfTasks
}

func GetTaskById(id string)(bool, models.Task){
	for _, task := range listOfTasks{
		if task.ID == id{
			return true, task
		}
	}
	return false, models.Task{}
}

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
func CreateTask(task models.Task){
	listOfTasks = append(listOfTasks, task)
}

func UpdateTask(id string, task models.Task)bool{
	found := false
	for i := 0; i < len(listOfTasks); i++{
		if listOfTasks[i].ID == id{
			listOfTasks[i] = task
			found = true
			fmt.Println("no")
			break
		}
	}
	return found
}