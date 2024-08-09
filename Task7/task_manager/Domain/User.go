package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionUser string = "user"

type User struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role string `json:"role" bson:"role"`
}

type UserUseCase interface{
	Register(user User) (interface{}, error)
	Login(user User) (string, error)
}

type UserRepository interface{
	Register(user User)(interface{}, error)
	Login(user User)(string, error)
}