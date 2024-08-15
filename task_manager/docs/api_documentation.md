
[Task Management API Documentation](https://documenter.getpostman.com/view/37482165/2sA3s7ioMo)

# Task Management API Documentation

## Overview

The Task Management API is a simple RESTful API that allows you to manage tasks. You can create, read, update, and delete tasks using this API. The API is built using Go and the Gin framework and uses an in-memory database for data storage.

## Base URL

The base URL for all endpoints is: http://localhost:8080

## Endpoints

### 1. Get All Tasks

**Endpoint:** `GET /tasks`

**Description:** Retrieve a list of all tasks.

**Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  [
    {
      "id": 1,
      "title": "Sample Task",
      "description": "This is a sample task",
      "due_date": "2024-08-06",
      "status": "pending"
    },
    {
      "id": 2,
      "title": "Another Task",
      "description": "This is another task",
      "due_date": "2024-08-06",
      "status": "completed"
    },
    {
      "id": 3,
      "title": "Yet Another Task",
      "description": "This is yet another task",
      "due_date": "2024-08-06",
      "status": "pending"
    }
  ]
  ```

### 2. Get Task by ID

**Endpoint:** GET /tasks/:id

**Description:** Retrieve the details of a specific task by its ID.

**Path Parameters:** `id` (integer): The ID of the task.

**Response:**

- **Status Code:** 200 OK
- **Body:**
    ```json
    {
        "id": 1,
        "title": "Sample Task",
        "description": "This is a sample task",
        "due_date": "2024-08-06",
        "status": "pending"
    }
    ```

### 3. Create a New Task

**Endpoint:** POST /tasks

**Description:** This endpoint allows you to create a new task. The request body should contain the details of the task to be created. The `title` field is required, while the `description`, `due_date`, and `status` fields are optional.

**Request:** 

- **Body:**
    ```json
    {
        "title": "New Task",
        "description": "This is a new task",
        "due_date": "2024-08-06",
        "status": "pending"
    }
    ```
- **Response:**

- **Status Code:** `201 Created`
- **Body:**
  ```json
  {
    "message": "Task created successfully",
    "task": {
      "id": 4,
      "title": "New Task",
      "description": "This is a new task",
      "due_date": "2024-08-06",
      "status": "pending"
    }
  }
  ```

### 4. Update a Task

**Endpoint:** PUT /tasks/:id

**Description:** Update the details of a specific task. The request body should contain the updated details of the task. The `title` field is required, while the `description`, `due_date`, and `status` fields are optional.

- **Path Parameters:** `id` (integer): The ID of the task.

- **Request:**
    ```json
    {
        "title": "Updated Task",
        "description": "This is an updated task",
        "due_date": "2024-08-06",
        "status": "completed"
    }
    ```
- **Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  {
    "message": "Task updated successfully",
    "task": {
      "id": 1,
      "title": "Updated Task",
      "description": "This is an updated task",
      "due_date": "2024-08-06",
      "status": "completed"
    }
  }
  ```

### 5. Delete a Task

**Endpoint:** DELETE /tasks/:id

**Description:** This endpoint allows you to delete a specific task by its ID.

- **Path Parameters:** `id` (integer): The ID of the task.

- **Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```
