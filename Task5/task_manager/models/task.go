package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* A struct to define the task*/
type Task struct {
	ID          primitive.ObjectID  `bson:"_id"`
	Title       string  `bson:"title"`
	Description string  `bson:"description"`
	DueDate     time.Time `bson:"dueDate"`
	Status      string  `bson:"status"`
}
