package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	PasswordHash string    `json:"-"`
	Role         string    `gorm:"not null" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

type Course struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `json:"name"`
	TeacherID uuid.UUID `gorm:"type:uuid;not null"`
	Teacher   User      `gorm:"foreignKey:TeacherID"`
	Credits   float64   `json:"credits"`
	CreatedAt time.Time `json:"created_at"`
}

type Grade struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	StudentID uuid.UUID `gorm:"type:uuid;not null"`
	Student   User      `gorm:"foreignKey:StudentID"`
	CourseID  uuid.UUID `gorm:"type:uuid;not null"`
	Course    Course    `gorm:"foreignKey:CourseID"`
	Grade     float64   `json:"grade"`
	CreatedAt time.Time `json:"created_at"`
}