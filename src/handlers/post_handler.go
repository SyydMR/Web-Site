package handlers

import (
	"net/http"
	"strconv"

	"github.com/SyydMR/Web-Site/src/models"
	"github.com/gin-gonic/gin"
)


func HandlerGetAllPosts (c *gin.Context) {
	posts, err := models.GetAllPost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}



func HandlerGetPostByID(c *gin.Context) {
	postID := c.Param("postId")
	id, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := models.GetPostById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func HandlerCreateEmptyPost(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	post := models.Post{
		AuthorID: id, 
		Contents: []models.Content{}, 
		Publish:  false,
	}

	err = models.CreatePost(&post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/blogs/"+strconv.FormatUint(uint64(post.ID), 10)+"/update")
}

func HandlerGetUserAllPosts (c *gin.Context) {
	id, err := getUserId(c)
	posts, err := models.GetPostsByUserID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}



func HandlerDeletePost (c *gin.Context) {
	postID := c.Param("postId")
	id, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    if err := models.DeletePost(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}