package handlers

import (
	"github.com/SyydMR/Web-Site/src/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	newUsers, err := models.GetAllUser()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return
	}
	c.JSON(http.StatusOK, newUsers)
}

func GetUserByIDHandler(c *gin.Context) {
	userId := c.Param("userId")

	ID, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	newUser, err := models.GetUserById(uint(ID))
	if err != nil {
		log.Println("Error retrieving user by ID:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, newUser)
}




func DeleteUserHandler(c *gin.Context) {
	userId := c.Param("userId")

    ID, err := strconv.ParseUint(userId, 10, 64)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    err = models.DeleteUser(uint(ID))
    if err!= nil {
        log.Printf("Error deleting user with ID %d: %v", ID, err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


func DeleteAllUsers(c *gin.Context) {
	err := models.DeleteAllUsers()
    if err!= nil {
        log.Printf("Error deleting all users: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all users"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "All users deleted successfully"})
}