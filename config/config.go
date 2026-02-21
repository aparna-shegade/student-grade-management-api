package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "grade-api/models"  // Add this import
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=1234 dbname=gradedb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect database")
    }
    DB = db
    db.AutoMigrate(&models.User{}, &models.Course{}, &models.Grade{})
}
