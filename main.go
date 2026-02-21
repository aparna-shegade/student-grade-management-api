// @title Student Grade Management API
// @version 1.0
// @description University-style grade management system
// @host localhost:8080
// @BasePath /

package main

import (
	"grade-api/config"
	"grade-api/controllers"
	"grade-api/middleware"
	"grade-api/seed"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "grade-api/docs" // VERY IMPORTANT
)

func main() {
	config.ConnectDatabase()
	seed.SeedData()

	r := gin.Default()

	// Swagger route MUST come before r.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/login", controllers.Login)

	// Protected routes
	api := r.Group("/api").Use(middleware.Auth())
	{
		api.POST("/courses", middleware.TeacherOrAdmin(), controllers.CreateCourse)
		api.GET("/courses", controllers.ListCourses)
		api.POST("/grades", middleware.TeacherOrAdmin(), controllers.AddGrade)
		api.GET("/my-grades", controllers.GetMyGrades)
		api.GET("/gpa", controllers.GetGPA)
	}

	r.Run(":8080")
}