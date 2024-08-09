package repositories

import (
	"context"
	"errors"
	"log"
	"strings"
	domain "task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(dbCollection *mongo.Collection)*TaskRepository{
	return &TaskRepository{
		collection: dbCollection,
	}
}

func (taskRepository *TaskRepository)GetAllTasks(role string, uid string)([]domain.Task, error){
	var filter bson.D
	var tasks []domain.Task
	var cursor *mongo.Cursor
	option := options.Find()
	user_id, err := primitive.ObjectIDFromHex(uid)
	if err != nil && strings.ToUpper(role) == "USER"{
		log.Fatal(err)
		return nil, errors.New("Invalid user id.")
	}
	if strings.ToUpper(role) == "ADMIN"{
		filter = bson.D{{}}
	}else{
		filter = bson.D{{Key:"uid", Value: user_id}}
	}
	cursor, err = taskRepository.collection.Find(context.TODO(), filter, option)
	if err != nil{
		log.Fatal(err)
		return nil, errors.New("Task not found.")
	}
	for cursor.Next(context.TODO()){
		var task domain.Task
		err := cursor.Decode(&task)
		if err != nil{
			log.Fatal(err)
			return nil, errors.New("Task not found.")
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func (taskRepository *TaskRepository) GetTaskById(id string, uid string, role string)(domain.Task, error){
	var task domain.Task
	task_id, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		//log.Fatal(err)
		return domain.Task{}, errors.New("Task not found")
	}
	user_id, err := primitive.ObjectIDFromHex(uid)
	if err != nil && strings.ToUpper(role) == "USER"{
		//log.Fatal(err)
		return domain.Task{}, errors.New("user not logged in")
	}
	err = taskRepository.collection.FindOne(context.TODO(),bson.D{{Key: "_id", Value: task_id}} ).Decode(&task)
	if err != nil{
		//log.Fatal(err)
		return domain.Task{}, errors.New("task Not found")
	}
	if strings.ToUpper(role) == "USER" && task.UserId != user_id{
		return domain.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (taskRepository *TaskRepository) DeleteTaskById(id string, uid string, role string)(bool, error){
	var filter bson.D
	task_id, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		//log.Fatal(err)
		return false, errors.New("task not found")
	}
	user_id, err := primitive.ObjectIDFromHex(uid)
	if err != nil && strings.ToUpper(role) == "USER"{
		//log.Fatal(err)
		return false, errors.New("user not logged in")
	}
	if strings.ToUpper(role) == "ADMIN"{
		filter = bson.D{{Key: "_id", Value: task_id}}
	}else{
		filter = bson.D{{Key: "_id", Value: task_id}, {Key: "uid", Value: user_id}}
	}
	result, err := taskRepository.collection.DeleteOne(context.TODO(), filter)
	if err != nil{
		//log.Fatal(err)
		return false, errors.New("task not found")
	}
	return result.DeletedCount > 0, nil
}

func (taskRepository *TaskRepository) CreateTask(newTask domain.Task, uid string, role string)(interface{}, error){
	user_id, err := primitive.ObjectIDFromHex(uid)
	if err != nil && strings.ToUpper(role) == "USER"{
		log.Fatal(err)
		return nil, errors.New("user not logged in")
	}
	newTask.ID = primitive.NewObjectID()
	newTask.UserId = user_id
	task, err := bson.Marshal(newTask)
	if err != nil{
		log.Fatal(err)
		return nil, errors.New("task not created")
	}
	result, err := taskRepository.collection.InsertOne(context.TODO(),task)
	if err != nil{
		log.Fatal(err)
		return nil, errors.New("task not created")
	}
	return result.InsertedID, nil
}

func (taskRepository *TaskRepository) UpdateTask(id string, uid string, role string, newTask domain.Task)(bool, error){
	var updatedTask bson.D
	var filter bson.D
	task_id, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		// log.Fatal(err)
		return false, errors.New("task not found")
	}
	user_id, err := primitive.ObjectIDFromHex(uid)
	if err != nil && strings.ToUpper(role) == "USER"{
		log.Fatal(err)
		return false, errors.New("User not logged in")
	}
	if newTask.Title != ""{
		updatedTask = append(updatedTask, bson.E{Key: "title", Value : newTask.Title})
	}
	if newTask.Description != ""{
		updatedTask = append(updatedTask, bson.E{Key: "description", Value : newTask.Description})
	}
	if newTask.Status != ""{
		updatedTask = append(updatedTask, bson.E{Key: "status", Value : newTask.Status})
	}
	// if !newTask.DueDate.IsZero(){
	// 	updatedTask = append(updatedTask, bson.E{Key: "dueDate", Value : newTask.DueDate})
	// }
	if strings.ToUpper(role) == "ADMIN"{
		filter = bson.D{{Key: "_id", Value: task_id}}
	}else{
		filter = bson.D{{Key: "_id", Value: task_id}, {Key: "uid", Value: user_id}}
	}
	result, err := taskRepository.collection.UpdateOne(context.TODO(), filter, bson.D{{Key : "$set", Value: updatedTask}})
	if err != nil{
		// log.Fatal(err)
		return false, errors.New("task not updated")
	}
	return result.ModifiedCount > 0, nil
}
