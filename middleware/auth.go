package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret-key-change-in-prod")

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
            c.Abort()
            return
        }
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
            c.Abort()
            return
        }
        tokenStr := parts[1]
        token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })
        if token == nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        claims := token.Claims.(jwt.MapClaims)
        c.Set("role", claims["role"])
        c.Set("user_id", claims["user_id"])
        c.Next()
    }
}

func TeacherOrAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, _ := c.Get("role")
        if role != "teacher" && role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Teacher/Admin only"})
            c.Abort()
            return
        }
        c.Next()
    }
}
