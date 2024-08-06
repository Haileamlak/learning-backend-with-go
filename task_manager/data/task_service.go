package data

import (
	"errors"
	"task_manager/models"
)

var tasks []models.Task = []models.Task{
	{ID: 1, Title: "Task 1", DueDate: "10/5/2024", Description: "Description 1", Status: "In Progress"},
	{ID: 2, Title: "Task 2", DueDate: "4/6/2024", Description: "Description 2", Status: "Completed"},
	{ID: 3, Title: "Task 3", DueDate: "4/4/2024", Description: "Description 3", Status: "Not Started"},
	{ID: 4, Title: "Task 4", DueDate: "3/7/2024", Description: "Description 4", Status: "In Progress"},
	{ID: 5, Title: "Task 5", DueDate: "1/8/2024", Description: "Description 5",	Status: "Not Started"},
}

var nextID int = 6

func GetTasks() []models.Task {
	return tasks
}

func GetTaskByID(id int) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func CreateTask(task models.Task) models.Task {
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id int, newTask models.Task) (models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {

			if newTask.Title != "" {
				tasks[i].Title = newTask.Title
			}
			if newTask.Description != "" {
				tasks[i].Description = newTask.Description
			}
			if newTask.DueDate != "" {
				tasks[i].DueDate = newTask.DueDate
			}
			if newTask.Status != "" {
				tasks[i].Status = newTask.Status
			}

			return tasks[i], nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
