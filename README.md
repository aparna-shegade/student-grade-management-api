# Student Grade Management API

## Roles

| Role    | Create Course | Add Grade | View GPA |
|---------|--------------|----------|----------|
| Admin   | Yes          | Yes      | Yes      |
| Teacher | Yes          | Yes      | No       |
| Student | No           | No       | Yes      |

## Setup

1. Install Go
2. Install PostgreSQL
3. Create database gradedb
4. Update config.go with credentials
5. Run:

go run main.go

## Seed Users

admin1 / password123
teacher1 / password123
student1 / password123