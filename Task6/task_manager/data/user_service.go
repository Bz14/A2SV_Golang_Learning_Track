package data

import (
	"context"
	"errors"
	"task_manager/database"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	userCollection *mongo.Collection
}
var jwtSecret = []byte("secret")

func CreateUserCollection()*mongo.Collection{
	connection := database.ConnectToDatabase()
	collection := connection.Database("task_manager").Collection("users")
	return collection
}

func UserCollection()*Users{
	collection := CreateUserCollection()
	return &Users{
		userCollection: collection,
	}
}

func (users *Users) Register(user models.User)(interface{}, error){
	var existingUser models.User
	err := users.userCollection.FindOne(context.TODO(), bson.D{{Key: "username", Value : user.UserName}}).Decode(&existingUser)
	if err == nil{
		return nil, errors.New("username already in use")
	}
	hashedPassword := hashPassword(user.Password)
	if hashedPassword == nil{
		return nil, errors.New("can not create password")
	}
	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()
	newUser , err := bson.Marshal(user)
	if err != nil{
		return nil, errors.New("can not create user")
	}
	createdId , err := users.userCollection.InsertOne(context.TODO(), newUser)
	if err != nil{
		return nil, errors.New("can not create user")
	}
	return createdId.InsertedID, nil
}

func (users *Users) Login(user models.User)(string, error){
	var loggedUser models.User
	err := users.userCollection.FindOne(context.TODO(), bson.D{{Key : "username", Value : user.UserName}}).Decode(&loggedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("invalid username or password")
		}
	}
	if !unHashPassword([]byte(loggedUser.Password), []byte(user.Password)) {
		return "", errors.New("invalid username or password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id" : loggedUser.ID,
		"username" : loggedUser.UserName,
		"role" : loggedUser.Role,
		"expires" : (time.Now().Add(5 * time.Minute)).Unix(),
	})
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil{
		return "", errors.New("user not logged in")
	}
	return jwtToken, nil
}

func hashPassword(password string)[]byte{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	return hashedPassword
}

func unHashPassword(existingPassword []byte, newPassword []byte)bool{
	return bcrypt.CompareHashAndPassword(existingPassword, newPassword) == nil
}