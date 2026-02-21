package controllers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "grade-api/config"
    "grade-api/models"
    "grade-api/services"
)

func CreateCourse(c *gin.Context) {
    var course models.Course
    if err := c.BindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    course.ID = uuid.New()
    course.CreatedAt = time.Now()
    config.DB.Create(&course)
    c.JSON(http.StatusOK, course)
}

func ListCourses(c *gin.Context) {
    var courses []models.Course
    config.DB.Find(&courses)
    c.JSON(http.StatusOK, courses)
}

func AddGrade(c *gin.Context) {
    var input struct {
        StudentID uuid.UUID `json:"student_id"`
        CourseID  uuid.UUID `json:"course_id"`
        Grade     float64   `json:"grade"`
    }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    grade := models.Grade{
        ID:        uuid.New(),
        StudentID: input.StudentID,
        CourseID:  input.CourseID,
        Grade:     input.Grade,
        CreatedAt: time.Now(),
    }
    config.DB.Create(&grade)
    c.JSON(http.StatusOK, grade)
}

func GetMyGrades(c *gin.Context) {
    userIDStr := c.GetString("user_id")
    userID, _ := uuid.Parse(userIDStr)
    grades := services.GetStudentGrades(userID)
    c.JSON(http.StatusOK, grades)
}

func GetGPA(c *gin.Context) {
    userIDStr := c.GetString("user_id")
    userID, _ := uuid.Parse(userIDStr)
    gpa := services.CalculateGPA(userID)
    c.JSON(http.StatusOK, gin.H{"gpa": gpa, "scale": "4.0"})
}
