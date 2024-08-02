package models

import "time"

/* A struct to define the task*/
type Task struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Status      string  `json:"status"`
}
