package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionTask string = "tasks"
const CollectionUser string = "user"

type Task struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     string      `json:"dueDate" bson:"dueDate"`   //time.Time 
	Status      string             `json:"status" bson:"status"`
	UserId      primitive.ObjectID `json:"uid" bson:"uid"`
}

type User struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role string `json:"role" bson:"role"`
}

type TaskRepository interface{
	GetAllTasks(role string, uid string)([]Task, error)
	GetTaskById(id string, uid string, role string)(Task, error)
	DeleteTaskById(id string, uid string, role string)(bool, error)
	CreateTask(newTask Task, uid string, role string)(interface{}, error)
	UpdateTask(id string, uid string, role string, newTask Task)(bool, error)
}

type TaskUseCase interface{
	GetAllTasks(role string, uid string)([]Task, error)
	GetTaskById(id string, uid string, role string)(Task, error)
	DeleteTaskById(id string, uid string, role string)(bool, error)
	CreateTask(newTask Task, uid string, role string)(interface{}, error)
	UpdateTask(id string, uid string, role string, newTask Task)(bool, error)
}

type UserUseCase interface{
	Register(user User) (interface{}, error)
	Login(user User) (string, error)
}

type UserRepository interface{
	Register(user User)(interface{}, error)
	Login(user User)(string, error)
}