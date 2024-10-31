package utils

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	// "io"
	// "net/http"
	"strconv"
)

// func ParseBody(r *http.Request, x interface{}) error {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		return err
// 	}
// 	defer r.Body.Close()

// 	if err := json.Unmarshal(body, x); err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetUserId(c *gin.Context) (uint, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, fmt.Errorf("authorization token not provided")
	}
	userID, err := VerifyJWT(tokenString)
	if err != nil {
		return 0, fmt.Errorf("invalid token: %v", err)
	}
	return userID, nil
}

func GetPostIDParam(c *gin.Context) (uint, error) {
	postID := c.Param("postId")
	id, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid post ID")
	}
	return uint(id), nil
}

func GetTaskIDParam(c *gin.Context) (uint, error) {
	taskID := c.Param("taskID")
	id, err := strconv.ParseUint(taskID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid task ID")
	}
	return uint(id), nil
}
