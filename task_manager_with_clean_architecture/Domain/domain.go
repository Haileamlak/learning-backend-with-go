package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username" json:"username" required:"true" binding:"required"`
	Password string             `bson:"password" json:"password" required:"true" binding:"required"`
	Role     string             `bson:"role" json:"role"`
}

type Task struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title   string             `bson:"title" json:"title" required:"true" binding:"required"`
	DueDate time.Time          `bson:"due_date" json:"due_date" required:"true" binding:"required"`
	Status  string             `bson:"status" json:"status" required:"true" binding:"required"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}

	if t.DueDate.IsZero() {
		return errors.New("due date is required")
	}

	if t.Status == "" {
		return errors.New("status is required")
	}

	if t.Status != "pending" && t.Status != "completed" {
		return errors.New("status must be either pending or completed")
	}

	if t.Status == "completed" && time.Now().Before(t.DueDate) {
		return errors.New("due date must be in the past")
	}

	if t.Status == "pending" && time.Now().After(t.DueDate) {
		return errors.New("due date must be in the future")
	}

	return nil
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type UserAlreadyExistsError struct {
	Message string
}

func (e *UserAlreadyExistsError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

type ForbiddenError struct {
	Message string
}

func (e *ForbiddenError) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}