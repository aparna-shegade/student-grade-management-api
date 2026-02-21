# Student Grade Management API

A RESTful API built using Go (Gin framework) for managing student grades similar to university portals like Canvas or Blackboard.

---

## Features

- JWT-based authentication
- Role-based access control (Admin, Teacher, Student)
- Course creation and management
- Grade entry and management
- GPA calculation
- PostgreSQL database integration
- Swagger API documentation
- Seed data initialization

---

## Roles

| Role    | Register | Create Course | Add Grade | View Grades | View GPA |
|---------|----------|--------------|-----------|------------|----------|
| Admin   | Yes      | Yes          | Yes       | Yes        | Yes      |
| Teacher | No       | Yes          | Yes       | Yes        | No       |
| Student | No       | No           | No        | Yes        | Yes      |

---

## Setup

1. Install Go
2. Install PostgreSQL
3. Create database:

CREATE DATABASE gradedb;

4. Update config/config.go with your database credentials
5. Install dependencies:

go mod tidy

6. Run the application:

go run main.go

Server will start at:

http://localhost:8080

---

## API Documentation

Swagger UI available at:

http://localhost:8080/swagger/index.html

---

## Seed Users

admin1    / password123
teacher1  / password123
student1  / password123

---

## Project Structure

controllers/
models/
middleware/
seed/
config/
docs/
main.go
go.mod
go.sum
README.md
AI_PROMPTS.md
DESIGN.md

---

## Authentication

- Passwords are securely hashed using bcrypt
- JWT tokens are generated on login
- Protected routes require:

Authorization: Bearer <token>

---

## AI Assistance

AI tools were used for:
- JWT implementation guidance
- Middleware structure
- Swagger integration
- Debugging and issue resolution

All prompts used during development are documented in AI_PROMPTS.md.

---

## Author

Aparna Shegade  
Capstone Project – Go (Golang)