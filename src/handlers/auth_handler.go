package handlers

import (
	"net/http"

	"github.com/SyydMR/Web-Site/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}
func RegisterHandler(c *gin.Context) {
	createUser := &models.User{}

	if err := c.ShouldBindJSON(createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

    registeredUser, err := createUser.Register()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, registeredUser)
}

func LoginHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user := models.User{Username: loginData.Username}

	token, err := user.Login(loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
