package repositories

import (
	"context"
	"log"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TaskRepository interface
type TaskRepository interface {
	CreateTask(task domain.Task) error
	GetTask(id string) (domain.Task, error)
	GetTasks() ([]domain.Task, error)
	UpdateTask(id string, task domain.Task) error
	DeleteTask(id string) error
}

// taskRepository struct
type taskRepository struct {
	db *mongo.Database
}

// NewTaskRepository creates a new task repository
func NewTaskRepository() TaskRepository {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	database := client.Database("task_manager")
	return &taskRepository{db: database}
}

// CreateTask creates a new task
func (r *taskRepository) CreateTask(task domain.Task) error {
	task.ID = ""
	_, err := r.db.Collection("tasks").InsertOne(context.TODO(), task)

	if err != nil {
		return &domain.InternalServerError{Message: "Error creating task"}
	}
	return nil
}

// GetTask retrieves a task by ID
func (r *taskRepository) GetTask(id string) (domain.Task, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, &domain.BadRequestError{Message: "Invalid ID"}
	}

	filter := bson.M{"_id": objId}
	var task domain.Task
	err = r.db.Collection("tasks").FindOne(context.TODO(), filter).Decode(&task)

	if err == mongo.ErrNoDocuments {
		return domain.Task{}, &domain.NotFoundError{Message: "Task not found"}
	}

	if err != nil {
		return domain.Task{}, &domain.InternalServerError{Message: "Error retriving task"}
	}

	return task, nil
}

// GetTasks retrieves all tasks
func (r *taskRepository) GetTasks() ([]domain.Task, error) {
	cursor, err := r.db.Collection("tasks").Find(context.TODO(), bson.M{})

	if err == mongo.ErrNoDocuments {
		return nil, &domain.NotFoundError{Message: "Tasks not found"}
	}

	if err != nil {
		return nil, &domain.InternalServerError{Message: "Error retrieving tasks"}
	}

	defer cursor.Close(context.TODO())

	var tasks []domain.Task
	for cursor.Next(context.TODO()) {
		var task domain.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// UpdateTask updates a task
func (r *taskRepository) UpdateTask(id string, task domain.Task) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.BadRequestError{Message: "Invalid ID"}
	}

	filter := bson.M{"_id": objId}

	update := bson.M{
		"$set": bson.M{
			"title":    task.Title,
			"due_date": task.DueDate,
			"status":   task.Status,
		},
	}

	_, err = r.db.Collection("tasks").UpdateOne(context.TODO(), filter, update)

	if err == mongo.ErrNoDocuments {
		return &domain.NotFoundError{Message: "Task not found"}
	}

	if err != nil {
		return &domain.InternalServerError{Message: "Error updating task"}
	}
	return nil
}

// DeleteTask deletes a task
func (r *taskRepository) DeleteTask(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.BadRequestError{Message: "Invalid ID"}
	}

	filter := bson.M{"_id": objId}

	_, err = r.db.Collection("tasks").DeleteOne(context.TODO(), filter)

	if err == mongo.ErrNoDocuments {
		return &domain.NotFoundError{Message: "Task not found"}
	}

	if err != nil {
		return &domain.InternalServerError{Message: "Error deleting task"}
	}

	return nil
}
