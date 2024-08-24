package handlers

import (
	"net/http"

	"github.com/ezekielnizamani/JobScam/internal/database"
	"github.com/ezekielnizamani/JobScam/internal/models"
	"github.com/ezekielnizamani/JobScam/internal/services"
	"github.com/gin-gonic/gin"
)

// @Summary Sign Up a new user
// @Description Create a new user with a username, email, and password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   user  body    models.User  true  "User data" // The user object containing username, email, and password
// @Success 201 {object} map[string]string "User created successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/auth/signup [post]
func SignUpHandler(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate required fields
	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, email, and password are required"})
		return
	}
	// Hash the password
	hashedPassword, err := services.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	newUser.Password = hashedPassword

	// Create the user
	if err := database.GetDB().Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while saving user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// @Summary Sign In a user
// @Description Authenticates a user with a username and password, and returns a JWT token if successful
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   user  body    models.User  true  "User credentials" // The user object containing username and password
// @Success 200 {object} map[string]string "Successful sign-in with JWT token"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/auth/signin [post]
func SignInHandler(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := database.GetDB().Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !services.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := services.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
