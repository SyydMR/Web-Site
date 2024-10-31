package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SyydMR/Web-Site/src/models"
	"github.com/SyydMR/Web-Site/src/utils"

	"github.com/gin-gonic/gin"

)

func getUserId(c *gin.Context) (uint, error) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        return 0, fmt.Errorf("authorization token not provided")
    }
    userID, err := utils.VerifyJWT(tokenString)
    if err != nil {
        return 0, fmt.Errorf("invalid token: %v", err)
    }
    return userID, nil
}

func GetAllTasks(c *gin.Context) {
	Id, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	Id = uint(Id)

	newTasks, err := models.GetUserAllTask(Id)

	c.JSON(http.StatusOK, newTasks)
}

func AddTask(c *gin.Context) {
	Id, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing task data"})
		return
	}

	Id = uint(Id)
	user, err := models.GetUserById(Id)

	newTask.UserID = Id
	err = user.AddTask(&newTask)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding task"})
		return
	}

	userTasks, err := models.GetUserAllTask(Id)
	if err != nil {
		log.Printf("Error fetching tasks for user %d: %v", Id, err)
		return
	}
	c.JSON(http.StatusOK, userTasks)
}

func RemoveTask(c *gin.Context) {
	taskID := c.Param("TaskId")
	ID, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	err = models.RemoveTask(uint(ID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove task"})
		return
	}

	Id, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newTasks, err := models.GetUserAllTask(Id)
	if err != nil {
		log.Printf("Error fetching tasks for user %d: %v", Id, err)
		return
	}
	c.JSON(http.StatusOK, newTasks)
}

func UpdateTask(c *gin.Context) {
	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing task data"})
		return
	}

	taskID := c.Param("TaskId")
	ID, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	taskDetails, _ := models.GetTaskById(uint(ID))

	if updateTask.Title != "" {
		taskDetails.Title = updateTask.Title
	}
	if updateTask.Description != "" {
		taskDetails.Description = updateTask.Description
	}
	if updateTask.Status != "" {
		taskDetails.Status = updateTask.Status
	}

	db.Save(&taskDetails)

	Id, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newTasks, err := models.GetUserAllTask(Id)
	if err != nil {
		log.Printf("Error fetching tasks for user %d: %v", Id, err)
		return
	}
	c.JSON(http.StatusOK, newTasks)
}
