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

// UserRepository interface
type UserRepository interface {
	CreateUser(user domain.User) error
	UpdateUser(id string, user domain.User) error
	FindByUsername(username string) (domain.User, error)
	CountUsers() (int64, error)
}

// userRepository struct
type userRepository struct {
	db *mongo.Database
}

// NewUserRepository creates a new user repository
func NewUserRepository() UserRepository {
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	database := client.Database("task_manager")
	return &userRepository{db: database}
}

// CreateUser creates a new user
func (r *userRepository) CreateUser(user domain.User)error {
	_, err := r.db.Collection("users").InsertOne(context.TODO(), user)

	if err != nil {
		return &domain.InternalServerError{Message: "Error creating user"}
	}

	return nil
}

func (r *userRepository) UpdateUser(id string, user domain.User)error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &domain.BadRequestError{Message: "Invalid ID"}
	}
	user.ID = ""
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": user}
	_, err = r.db.Collection("users").UpdateOne(context.TODO(), filter, update)

	if err == mongo.ErrNoDocuments {
		return &domain.NotFoundError{Message: "User not found"}
	}

	if err != nil {
		return &domain.InternalServerError{Message: "Error updating user"}
	}

	return nil
}

func (r *userRepository) FindByUsername(username string) (domain.User,error) {
	var user domain.User
	filter := bson.M{"username": username}
	err := r.db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return domain.User{}, &domain.NotFoundError{Message: "User not found"}
	}

	if err != nil {
		return domain.User{}, &domain.InternalServerError{Message: "Error retrieving user"}
	}

	return user, nil
}

func (r *userRepository) CountUsers() (int64,error) {
	count, err := r.db.Collection("users").CountDocuments(context.TODO(), bson.M{})

	if err != nil {
		return 0, &domain.InternalServerError{Message: "Error counting users"}
	}

	return count, nil
}
