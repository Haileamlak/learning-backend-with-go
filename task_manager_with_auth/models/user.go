package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Username string             `json:"username"`
    Password string             `json:"password"` // Stored as a hashed password
    Role     string             `json:"role"`     // "admin" or "user"
}