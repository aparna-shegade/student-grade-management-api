package services

import (
    "math"
    "grade-api/config"
    "grade-api/models"
    "github.com/google/uuid"
)

func CalculateGPA(studentID uuid.UUID) float64 {
    var grades []struct {
        Grade   float64
        Credits float64
    }
    config.DB.Table("grades g").
        Joins("JOIN courses c ON g.course_id = c.id").
        Where("g.student_id = ?", studentID).
        Select("g.grade / 100.0 * c.credits as grade, c.credits").
        Scan(&grades)
    
    var totalPoints, totalCredits float64
    for _, g := range grades {
        totalPoints += g.Grade * g.Credits
        totalCredits += g.Credits
    }
    if totalCredits == 0 {
        return 0
    }
    return math.Round(totalPoints/totalCredits*100) / 100
}

func GetStudentGrades(studentID uuid.UUID) []models.Grade {
    var grades []models.Grade
    config.DB.Where("student_id = ?", studentID).Find(&grades)
    return grades
}
