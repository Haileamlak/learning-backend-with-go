package data

import (
	"context"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService() *TaskService {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	collection := client.Database("task_manager").Collection("tasks")
	return &TaskService{collection: collection}
}

func (service *TaskService) CreateTask(task models.Task) (*mongo.InsertOneResult, error) {

	result, err := service.collection.InsertOne(context.TODO(), task)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *TaskService) GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	ctx := context.TODO()
	cursor, err := service.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var task models.Task
		err := cursor.Decode(&task)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	cursor.Close(ctx)

	return tasks, nil
}

func (service *TaskService) GetTaskByID(id primitive.ObjectID) (models.Task, error) {

	var result models.Task
	err := service.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)

	if err != nil {
		return models.Task{}, err
	}

	return result, nil
}

func (service *TaskService) UpdateTask(id primitive.ObjectID, newTask models.Task) (*mongo.UpdateResult, error) {

	update := bson.M{
		"$set": newTask,
	}
	result, err := service.collection.UpdateByID(context.TODO(), id, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *TaskService) DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error) {

	result, err := service.collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}