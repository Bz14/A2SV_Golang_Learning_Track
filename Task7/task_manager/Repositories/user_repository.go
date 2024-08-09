package repositories

import (
	"context"
	"errors"
	"log"
	domain "task_manager/Domain"
	infrastructure "task_manager/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	passwordInterface infrastructure.PasswordHash
	tokenInterface infrastructure.JWT
}

func NewUserRepository(dbCollection *mongo.Collection, passwordAlg infrastructure.PasswordHash, token infrastructure.JWT)*UserRepository{
	return &UserRepository{
		collection : dbCollection,
		passwordInterface : passwordAlg,
		tokenInterface: token,
	}
}

func (userRepository *UserRepository)Register(user domain.User)(interface{}, error){
	var existingUser domain.User
	err := userRepository.collection.FindOne(context.TODO(), bson.D{{Key : "username", Value : user.UserName}}).Decode(&existingUser)
	if err == nil{
		// log.Fatal(err)
		return nil, errors.New("username already in use")
	}
	hashedPassword := userRepository.passwordInterface.HashPassword(user.Password)
	if hashedPassword == nil{
		log.Fatal(err)
		return nil, errors.New("can not create password")
	}
	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	newUser, err := bson.Marshal(user)
	if err != nil{
		return nil, errors.New("can not create user")
	}
	createdUser, err := userRepository.collection.InsertOne(context.TODO(), newUser)
	if err != nil{
		return nil, errors.New("can not create user")
	}
	return createdUser.InsertedID, nil
}

func (userRepository *UserRepository) Login(user domain.User)(string, error){
	var loggedUser domain.User
	err := userRepository.collection.FindOne(context.TODO(), bson.D{{Key : "username", Value : user.UserName}}).Decode(&loggedUser)
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	err = userRepository.passwordInterface.UnHashPassword([]byte(loggedUser.Password), []byte(user.Password));
	if err != nil {
		log.Fatal(err)
		return "", errors.New("invalid username or password")
	}
    
	token, err := userRepository.tokenInterface.GenerateToken(loggedUser)
	if err != nil{
		return "", err
	}
	return token, nil
}