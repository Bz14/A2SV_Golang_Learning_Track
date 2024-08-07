package data

import (
	"context"
	"fmt"
	"log"
	"strings"
	"task_manager/models"

	"task_manager/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskManager struct{
	taskCollection *mongo.Collection
}

func CreateTaskCollection()*mongo.Collection{
	connection := database.ConnectToDatabase()
	collection := connection.Database("task_manager").Collection("tasks")
	return collection
}

func NewTaskManager() *TaskManager{
	collection := CreateTaskCollection()

	return &TaskManager{
		taskCollection: collection,
	}
}

/* List of all available tasks*/
func (t *TaskManager)GetAllTasks(role string, uid string)[]models.Task{
	option := options.Find()
	var tasks []models.Task
	var cursor *mongo.Cursor
	var err error
	var filter bson.D
	userId, _ := primitive.ObjectIDFromHex(uid)
	if strings.ToUpper(role) == "ADMIN"{
		filter = bson.D{{}}
	}else{
		filter = bson.D{{Key : "uid", Value : userId}}
	}
	cursor, err = t.taskCollection.Find(context.TODO(), filter, option)
	if err != nil{
		log.Fatal(err)
		return nil
	}
	for cursor.Next(context.TODO()){
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil{
			// log.Fatal(err)
			return nil
		}
		tasks = append(tasks, task)
	}
	return tasks
}

/* Get task by ID*/
func (t *TaskManager)GetTaskById(id string, uid string, role string)interface{}{
	var err error
	objId, _ := primitive.ObjectIDFromHex(id)
	userId, _ := primitive.ObjectIDFromHex(uid)
	var task models.Task
	err = t.taskCollection.FindOne(context.TODO(), bson.D{{Key : "_id", Value : objId}}).Decode(&task)
	if err != nil{
		return nil
	}
	fmt.Println("Task", task)
	if strings.ToUpper(role) == "USER" && task.UserId != userId{
		return nil
	}
	return task
}

/* Delete task with a given ID */
func (t *TaskManager) DeleteTaskById(id string, uid string, role string)int64{
	var err error
	var result *mongo.DeleteResult
	objId, _ := primitive.ObjectIDFromHex(id)
	user_id, _ := primitive.ObjectIDFromHex(uid)
	if strings.ToUpper(role) == "ADMIN"{
		result, err = t.taskCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id" , Value: objId}})
	}else{
		result, err = t.taskCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id" , Value: objId}, {Key : "uid", Value : user_id}})
	}
	if err != nil{
		return 0
	}
	return result.DeletedCount
}

/* Create a new task*/
func (t *TaskManager) CreateTask(newTask models.Task, user_id string)interface{}{
	newTask.ID = primitive.NewObjectID()
	userId, _ := primitive.ObjectIDFromHex(user_id)
	newTask.UserId = userId
	createdTask , err := bson.Marshal(newTask)
	if err != nil{
		return nil
	}
	task, err := t.taskCollection.InsertOne(context.TODO(), createdTask)
	if err != nil{
		return nil
	}
	return task.InsertedID
}

/* Updating a task with a given ID*/
func  (t *TaskManager) UpdateTask(id string, uid string, role string, newTask models.Task)(bool, error){
	objId, _ := primitive.ObjectIDFromHex(id)
	var filter primitive.D
	var update bson.D
	user_id, _ := primitive.ObjectIDFromHex(uid)
	
	if newTask.Title != ""{
		update = append(update, bson.E{Key: "title", Value: newTask.Title})
	}
	if newTask.Description != ""{
		update = append(update, bson.E{Key: "description", Value: newTask.Description})
	}
	if newTask.Status != ""{
		update = append(update, bson.E{Key: "status", Value: newTask.Status})
	}
	if newTask.DueDate != ""{
		update = append(update, bson.E{Key: "dueDate", Value: newTask.DueDate})
	}
	fmt.Println("--------------------------------")
	fmt.Println("Update", update)
	if strings.ToUpper(role) == "ADMIN"{
		filter = bson.D{{Key: "_id", Value : objId}}

	}else{
		filter = bson.D{{Key: "_id", Value : objId}, {Key: "uid", Value : user_id}}
	}
	fmt.Println("Filter", filter)
	result, err := t.taskCollection.UpdateOne(context.TODO(), filter, bson.D{{Key : "$set" , Value: update}})
	fmt.Println("Result", result)
	if err != nil{
		// log.Fatal(err)
		return false, err
	}
	fmt.Println(result.ModifiedCount)
	return result.ModifiedCount > 0, nil
}

