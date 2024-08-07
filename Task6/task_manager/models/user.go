package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"` // making it '-' does not bind correctly
	Role string `json:"role" bson:"role"`
}