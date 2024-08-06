# Task Management API Documentation

## Overview

The Task Management API allows you to manage tasks with basic CRUD (Create, Read, Update, Delete) operations. This API is built using Go and the Gin framework and uses an in-memory database for data storage.

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

**Description:** Create a new task.

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

**Description:** Update the details of a specific task.

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

**Description:** Delete a specific task by its ID.

- **Path Parameters:** `id` (integer): The ID of the task.

- **Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```
