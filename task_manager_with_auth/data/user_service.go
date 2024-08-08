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

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService() *UserService {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	collection := client.Database("task_manager").Collection("users")
	return &UserService{collection: collection}
}

func (service *UserService) CreateUser(user models.User) error {
	// check if user already exists
	var existingUser models.User
    err := service.collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
    if err == nil {
        return fmt.Errorf("username already exists")
    }

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// If the database is empty, the first created user will be an admin.
	count, err := service.collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	// Save user to database
	_, err = service.collection.InsertOne(context.TODO(), user)

	return err
}

func (service *UserService) AuthenticateUser(username, password string) (models.User, error) {
	var user models.User
	err := service.collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserService) PromoteUser(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	_, err := s.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
