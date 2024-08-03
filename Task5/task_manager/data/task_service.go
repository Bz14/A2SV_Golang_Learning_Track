package data

import (
	"context"
	"fmt"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskManager struct{
	Collection *mongo.Collection
}

func NewTaskManager() *TaskManager{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	connection, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	err = connection.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err)
	}
	collection := connection.Database("task_manager").Collection("tasks")
	fmt.Println("Connected")
	// defer connection.Disconnect(context.TODO())
	
	return &TaskManager{
		Collection: collection,
	}
}


/* List of all available tasks*/
func (t *TaskManager)GetAllTasks()interface{}{
	var tasks []models.Task
	option := options.Find()
	cursor, err := t.Collection.Find(context.TODO(), bson.D{{}}, option)
	if err != nil{
		log.Fatal(err)
		return nil
	}
	for cursor.Next(context.TODO()){
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil{
			log.Fatal(err)
			return nil
		}
		tasks = append(tasks, task)
	}
	return tasks
}

/* Get task by ID*/
func (t *TaskManager)GetTaskById(id  string)interface{}{
	objId, _ := primitive.ObjectIDFromHex(id)
	var task models.Task
	err := t.Collection.FindOne(context.TODO(), bson.D{{Key : "_id", Value : objId}}).Decode(&task)
	if err != nil{
		return nil
	}
	return task
}

/* Delete task with a given ID */
func (t *TaskManager) DeleteTaskById(id string)int64{
	objId, _ := primitive.ObjectIDFromHex(id)
	result, err := t.Collection.DeleteOne(context.TODO(), bson.D{{Key: "_id" , Value: objId}})
	if err != nil{
		log.Fatal(err)
		return 0
	}
	return result.DeletedCount
	
}

/* Create a new task*/
func (t *TaskManager) CreateTask(newTask models.Task)interface{}{
	newTask.ID = primitive.NewObjectID()
	createdTask , err := bson.Marshal(newTask)
	if err != nil{
		return nil
	}
	task, err := t.Collection.InsertOne(context.TODO(), createdTask)
	if err != nil{
		return nil
	}
	return task.InsertedID
}

/* Updating a task with a given ID*/
func  (t *TaskManager) UpdateTask(id string, newTask models.Task)bool{
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value : objId}}
	var update bson.D
	
	if newTask.Title != ""{
		update = append(update, bson.E{Key: "title", Value: newTask.Title})
	}
	if newTask.Description != ""{
		update = append(update, bson.E{Key: "description", Value: newTask.Description})
	}
	if newTask.Status != ""{
		update = append(update, bson.E{Key: "status", Value: newTask.Status})
	}
	if !newTask.DueDate.IsZero(){
		update = append(update, bson.E{Key: "dueDate", Value: newTask.DueDate})
	}
	
	result, err := t.Collection.UpdateOne(context.TODO(), filter, bson.D{{Key : "$set" , Value: update}})
	if err != nil{
		log.Fatal(err)
		return false
	}
	return result.ModifiedCount > 0
}

