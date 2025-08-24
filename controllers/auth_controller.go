package controllers

import (
	"books-api/config"
	"books-api/middlewares"
	"books-api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register a new user
// @Description Register with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body object true "User info" example({"username":"user123","password":"123456"})
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/register [post]
func Register(c *gin.Context) {
    var req models.User
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    _, err := config.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", req.Username, string(hashed))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

// Login godoc
// @Summary User login
// @Description Login with username and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body object true "Login info" example({"username":"user123","password":"123456"})
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/login [post]
func Login(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", req.Username).Scan(
		&user.ID, &user.Username, &user.Password,
	)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, _ := middlewares.GenerateToken(user.Username)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
