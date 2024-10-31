package middlewares

import (
	"github.com/SyydMR/Web-Site/src/models"
	"github.com/SyydMR/Web-Site/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func validateOwnership(c *gin.Context, resourceUserID, tokenUserID uint) bool {
	if resourceUserID != tokenUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
		return false
	}
	return true
}

func IDPostValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, err := utils.GetPostIDParam(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		post, err := models.GetPostById(postID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}

		userIdFromToken, err := utils.GetUserId(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !validateOwnership(c, post.AuthorID, userIdFromToken) {
			return
		}

		c.Next()
	}
}

func IDTaskValidateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID, err := utils.GetTaskIDParam(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		task, err := models.GetTaskById(taskID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}

		userIdFromToken, err := utils.GetUserId(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !validateOwnership(c, task.UserID, userIdFromToken) {
			return
		}

		c.Next()
	}
}
