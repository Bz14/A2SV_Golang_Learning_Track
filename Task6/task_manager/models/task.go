package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* A struct to define the task*/
type Task struct {
	ID          primitive.ObjectID  `json:"_id" bson:"_id"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	DueDate     string  `json:"dueDate" bson:"dueDate"`
	Status      string  `json:"status" bson:"status"`
	UserId     primitive.ObjectID  `json:"uid" bson:"uid"`
}