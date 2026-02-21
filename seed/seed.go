package seed

import (
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"grade-api/config"
	"grade-api/models"
)

func SeedData() {

	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	admin := models.User{
		ID: uuid.New(),
		Username: "admin1",
		PasswordHash: string(password),
		Role: "admin",
		CreatedAt: time.Now(),
	}

	teacher := models.User{
		ID: uuid.New(),
		Username: "teacher1",
		PasswordHash: string(password),
		Role: "teacher",
		CreatedAt: time.Now(),
	}

	student := models.User{
		ID: uuid.New(),
		Username: "student1",
		PasswordHash: string(password),
		Role: "student",
		CreatedAt: time.Now(),
	}

	config.DB.Create(&admin)
	config.DB.Create(&teacher)
	config.DB.Create(&student)

	course := models.Course{
		ID: uuid.New(),
		Name: "Distributed Systems",
		TeacherID: teacher.ID,
		Credits: 4,
		CreatedAt: time.Now(),
	}

	config.DB.Create(&course)

	grade := models.Grade{
		ID: uuid.New(),
		StudentID: student.ID,
		CourseID: course.ID,
		Grade: 8.5,
		CreatedAt: time.Now(),
	}

	config.DB.Create(&grade)
}