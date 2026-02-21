package controllers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "grade-api/config"
    "grade-api/models"
)
// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body object true "User credentials"
// @Success 200 {object} models.User
// @Router /auth/register [post]
func Register(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
        Role     string `json:"role"`
    }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
    user := models.User{
        ID:           uuid.New(),
        Username:     input.Username,
        PasswordHash: string(hash),
        Role:         input.Role,
        CreatedAt:    time.Now(),
    }
    config.DB.Create(&user)
    c.JSON(http.StatusOK, user)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body object true "Login credentials"
// @Success 200 {object} map[string]string
// @Router /auth/login [post]
func Login(c *gin.Context){
	println("LOGIN FUNCTION HIT")
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var user models.User
    result := config.DB.Where("username = ?", input.Username).First(&user)

    if result.Error != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    // 🔴 DEBUG LINE (temporary)
    println("Stored Hash:", user.PasswordHash)

    err := bcrypt.CompareHashAndPassword(
        []byte(user.PasswordHash),
        []byte(input.Password),
    )

    if err != nil {
        println("Password comparison failed")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    println("Password matched")

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID.String(),
        "role":    user.Role,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })

    tokenStr, _ := token.SignedString([]byte("secret-key-change-in-prod"))

    c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

