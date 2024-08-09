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

## Instructions for user registration.

To register a user, you need to send a POST request to the `/register` endpoint with the following JSON payload:

```json
{
  "username": "john_doe",
  "password": "password123"
}
```

This will create a new user with the specified username and password. The response will include the ID of the newly created user.

## Instructions for user login.

To login a user, you need to send a POST request to the `/login` endpoint with the following JSON payload:

```json
{
    "username": "john_doe",
    "password": "password123"
}
```

This will log in the user with the specified username and password. The response will include the Token of the logged-in user.

## Instructions for promoting a user to admin.

To promote a user to an admin, you need to send a PUT request to the `/promote/:id` endpoint with the ID of the user you want to promote.

This will promote the user with the specified ID to an admin.

## Login and JWT usage.

The API uses JWT (JSON Web Tokens) for user authentication. When a user logs in successfully, a JWT token is generated and returned in the response. This token can be used to authenticate the user for subsequent requests that require authentication. The token should be included in the `Authorization` header of the request with the `Bearer` scheme. For example:

```json
Authorization
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5fZG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.4S50
```

This token will be validated by the API to authenticate the user and authorize access to protected endpoints.

## Access control rules and endpoints.

The API has two types of users: regular users and admin users. Regular users can access the following endpoints: `/tasks`, `/tasks/:id`, `/register`, `/login`. Admin users can access all endpoints, including the `/promote/:id` endpoint to promote other users to admin. The API uses JWT tokens to authenticate and authorize users based on their role. Regular users are assigned the role `user`, while admin users are assigned the role `admin`. The API checks the role of the user from the JWT token and allows or denies access to endpoints based on the user's role. The API uses middleware to enforce access control rules and validate JWT tokens for authentication.

## Usage of protected endpoints.

To access protected endpoints that require authentication, you need to include the JWT token in the `Authorization` header of the request with the `Bearer` scheme. The API will validate the token and authenticate the user based on the token. If the token is valid and the user is authorized to access the endpoint, the API will process the request and return the response. If the token is invalid or the user is not authorized to access the endpoint, the API will return an error response with the appropriate status code.

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
6. **Register a User**: `POST /register`
7. **Login a User**: `POST /login`
8. **Promote a User**: `PUT /promote/:id`

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
    "Task Id": "66b4b5720ad4ec403d1e7bd8"
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
    "output": {
        "MatchedCount": 1,
        "ModifiedCount": 1,
        "UpsertedCount": 0,
        "UpsertedID": null
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
    "message": "Task deleted successfully",
    "output": {
        "DeletedCount": 1
    }
  }
  ```

### 6. Register a User

**Endpoint:** POST /register

**Description:** Register a new user.

- **Request:**
  ```json
  {
    "username": "john_doe",
    "password": "password123"
  }
  ```
- **Response:**

- **Status Code:** `201 Created`
- **Body:**

  ```json
  {
    "message": "User registered successfully",
  }
  ```
### 7. Login a User

**Endpoint:** POST /login

**Description:** Login a user.

- **Request:**
  ```json
  {
    "username": "john_doe",
    "password": "password123"
  }
  ```
- **Response:**

- **Status Code:** `200 OK`
- **Body:**

  ```json
  {
    "message": "User logged in successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5fZG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.4S50"
  }
  ```
### 8. Promote a User

**Endpoint:** PUT /promote

**Description:** Promote a user to an admin.

- **Request:**
  ```json
  {
    "username": "john_doe",
  }
  ```
- **Response:**

- **Status Code:** `200 OK`
- **Body:**
  ```json
  {
    "message": "User promoted to admin successfully"
  }
  ```

## Conclusion

This API provides a simple and easy-to-use interface for managing tasks. You can use the provided endpoints to create, read, update, and delete tasks, as well as register and login users. The API is built using Go and MongoDB, making it fast and efficient for handling task management operations.