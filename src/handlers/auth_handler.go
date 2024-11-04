package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SyydMR/Web-Site/src/internal/email"
	"github.com/SyydMR/Web-Site/src/models"
	"github.com/SyydMR/Web-Site/src/utils"
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
	fmt.Println(createUser.Email)
    welcomeMessage := fmt.Sprintf("<h1>\u062e\u0648\u0634 \u0622\u0645\u062f\u06cc\u062f\u060c %s!</h1><p>\u0628\u0647 \u0648\u0628\u200c\u0633\u0627\u06cc\u062a \u0645\u0627 \u062e\u0648\u0634 \u0622\u0645\u062f\u06cc\u062f. \u0627\u0632 \u062b\u0628\u062a\u200c\u0646\u0627\u0645 \u0634\u0645\u0627 \u0628\u0633\u06cc\u0627\u0631 \u062e\u0648\u0634\u062d\u0627\u0644\u06cc\u0645 \u0648 \u0627\u0645\u06cc\u062f\u0648\u0627\u0631\u06cc\u0645 \u062a\u062c\u0631\u0628\u0647 \u062e\u0648\u0628\u06cc \u062f\u0627\u0634\u062a\u0647 \u0628\u0627\u0634\u06cc\u062f.</p>", createUser.Username)

	err = email.EmailSender(
		"mohammadrezadaryabaki@gmail.com",
		[]string{createUser.Email},
		"Welcome to Web-Site!",
		welcomeMessage,
	)
	
	if err != nil {
        log.Fatal("Failed to send email:", err)
    }
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

func LogoutHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token not provided"})
		c.Abort()
		return
	}

	utils.ExpireToken(tokenString)
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
