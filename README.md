# Task Management Service

Backend microservice built in Go with MongoDB and RabbitMQ.

## Features

- CRUD for tasks
- Pagination & filtering
- MongoDB persistence
- Clean project structure

# Problem Breakdown & Design Decisions

Users need to manage tasks with basic CRUD functionality.

MongoDB was chosen as the primary database for its flexibility and easy document-based structure.

Each layer (handler, service, repository, and db) was separated to follow the Single Responsibility Principle.

The system is built as a microservice with clean interfaces, making it easy to plug into a larger ecosystem.

Gorilla Mux is used for HTTP routing due to its popularity and ease of use.

The design allows for easy expansion into additional services (like a user service) using REST, gRPC, or messaging (RabbitMQ).

Pagination and filtering are handled via query parameters for scalability.

# How to Run the Service

- Clone the repository: git clone git@github.com:abhi00999/task-management.git
- cd task-management
- Start your local MongoDB server (e.g., mongod or MongoDB Atlas)

Install dependencies:
- go mod tidy

Run the server:
- go run ./cmd/main.go
- The service will be available at http://localhost:8080

# 📖 API Documentation

Base URL: http://localhost:8080

Health Check
- GET /health

Response: "healthy"

Create Task
- POST /tasks

Request Body:
{
  "title": "Write unit tests",
  "status": "Pending"
}

Response:
{
  "id": "654fc0e0f1fbb01234567890",
  "title": "Write unit tests",
  "status": "Pending"
}

Get Tasks
- GET /tasks?status=Completed&limit=10&skip=0

Response:
[
  {
    "id": "654fc0e0f1fbb01234567890",
    "title": "Write tests",
    "status": "Completed"
  }
]

Update Task
- PUT /tasks/{id}

Request Body:
{
  "title": "Write tests (updated)",
  "status": "Completed"
}

Response: 200 OK

Delete Task
- DELETE /tasks/{id}
- Response: 204 No Content

# Microservice Concepts Demonstrated
Single Responsibility Principle:
- Handlers deal with HTTP layer.
- Services handle business logic.
- Repositories interact with the DB.
- DB setup is isolated.

Loose Coupling:
- Components are modular, making unit testing and extension easier.

Scalability:
- Stateless design allows for horizontal scaling.

Inter-Service Communication:
- REST is currently used.
- Can be extended to use gRPC or RabbitMQ for async communication.

Pagination and Filtering:
- GET /tasks supports query parameters for status, limit, and skip to support large data sets.
