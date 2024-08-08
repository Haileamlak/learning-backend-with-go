# Task Management API Documentation

## Overview

The Task Management API allows you to manage tasks with basic CRUD (Create, Read, Update, Delete) operations. This API is built using Go and the Gin framework and uses MongoDB for persistent data storage.

## Installation

To install the API, you need to have Go installed on your system. You can download and install Go from the official website.

After installing Go, you can clone the repository and install the dependencies by running the following commands in your terminal:

```bash
git clone github.com/haileamlak/learning-backend-with-go.git
cd task_manager_with_mongodb
go mod download
```

This will clone the repository and install the required dependencies for the API.


## Database Configuration

To run this API in you computer, you need to install the community edition of MongoDB on your system. You can download MongoDB from the official website and follow the installation     instructions for your operating system.

After installing MongoDB, you need to start the MongoDB server by running the following command in your terminal:

```bash
mongod
```

This will start the MongoDB server on the default port `27017`. You can now connect to the MongoDB server using the `mongo` shell or a MongoDB client like Compass.

The API uses a database named `task_manager` and a collection named `tasks` to store the task data. You can create the database and collection by running the following commands in the `mongo` shell:

```bash
use task_manager
db.createCollection("tasks")
```

This will create the `task_manager` database and the `tasks` collection in MongoDB.

## Running the API

To run the API, you need to have Go and MongoDB installed on your system. You can start the API by running the following command in the root directory of the project:

```bash
go run main.go
```

This will start the API server on `http://localhost:8080`.

## API Endpoints

The API provides the following endpoints for managing tasks:

1. **Get All Tasks**: `GET /tasks`
2. **Get Task by ID**: `GET /tasks/:id`
3. **Create a New Task**: `POST /tasks`
4. **Update a Task**: `PUT /tasks/:id`
5. **Delete a Task**: `DELETE /tasks/:id`

The details of each endpoint, including the request and response formats, are provided below.

### 1. Get All Tasks

**Endpoint:** `GET /tasks`

**Description:** Retrieve a list of all tasks.

**Response:**

- **Status Code:** `200 OK`
- **Body:**

  ```json
  [
    {
      "id": "64d0a6f9f8c30f7e6b2b2b2b",
      "title": "Sample Task",
      "description": "This is a sample task",
      "due_date": "2024-08-06",
      "status": "pending"
    },
    {
      "id": "64d0a6f9f8c30f7e6b2b2b2c",
      "title": "Another Task",
      "description": "This is another task",
      "due_date": "2024-08-06",
      "status": "completed"
    }
  ]
  ```

### 2. Get Task by ID

**Endpoint:** GET /tasks/:id

**Description:** Retrieve the details of a specific task by its ID.

**Path Parameters:** `id` (string): The ID of the task.

**Response:**

- **Status Code:** 200 OK
- **Body:**
  ```json
  {
    "id": "64d0a6f9f8c30f7e6b2b2b2b",
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
      "id": "64d0a6f9f8c30f7e6b2b2b2d",
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
      "id": "64d0a6f9f8c30f7e6b2b2b2b",
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

- **Path Parameters:** `id` (string): The ID of the task.

- **Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```