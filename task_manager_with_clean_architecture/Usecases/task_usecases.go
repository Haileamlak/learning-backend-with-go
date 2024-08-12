package usecases

import (
	domain "task-manager/Domain"
	repositories "task-manager/Repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskUsecase interface
type TaskUsecase interface {
	CreateTask(task domain.Task) error
	GetTask(id primitive.ObjectID) (domain.Task, error)
	GetTasks() ([]domain.Task, error)
	UpdateTask(id primitive.ObjectID, task domain.Task) error
	DeleteTask(id primitive.ObjectID) error
}

// taskUsecase struct
type taskUsecase struct {
	taskRepo repositories.TaskRepository
}

// NewTaskUsecase creates a new task usecase
func NewTaskUsecase(taskRepo repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo}
}

// CreateTask creates a new task
func (u *taskUsecase) CreateTask(task domain.Task) error {
	if err := task.Validate(); err != nil {
		return &domain.BadRequestError{Message: err.Error()}
	}

	// check if task already exists
	tasks, _ := u.taskRepo.GetTasks()
	for _, t := range tasks {
		if t.Title == task.Title {
			return &domain.BadRequestError{Message: "Task already exists"}
		}
	}

	return u.taskRepo.CreateTask(task)
}

// GetTask retrieves a task by ID
func (u *taskUsecase) GetTask(id primitive.ObjectID) (domain.Task, error) {
	return u.taskRepo.GetTask(id)
}

// GetTasks retrieves all tasks
func (u *taskUsecase) GetTasks() ([]domain.Task, error) {
	return u.taskRepo.GetTasks()
}

// UpdateTask updates a task
func (u *taskUsecase) UpdateTask(id primitive.ObjectID, task domain.Task) error {
	if err := task.Validate(); err != nil {
		return &domain.BadRequestError{Message: err.Error()}
	}

	return u.taskRepo.UpdateTask(id, task)
}

// DeleteTask deletes a task
func (u *taskUsecase) DeleteTask(id primitive.ObjectID) error {
	return u.taskRepo.DeleteTask(id)
}
