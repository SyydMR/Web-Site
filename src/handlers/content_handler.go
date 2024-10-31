package handlers

import (
	"net/http"
	"strconv"

	"github.com/SyydMR/Web-Site/src/models"
	"github.com/gin-gonic/gin"
)



func HandlerGetAllContent(c *gin.Context) {
	postID := c.Param("postId")
	id, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

    contents, err := models.GetAllContentByPostID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, contents)
}

func HandlerCreateContent(c *gin.Context) {
	postID := c.Param("postId")
	id, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var content models.Content

	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing content data"})
		return
	}
	
	newContent := &models.Content{
		PostID: uint(id),
		Type: content.Type,
		Data: content.Data,
	}

	post, err := models.GetPostById(uint(id))	
	post.CreateContent(*newContent)

	c.Redirect(http.StatusSeeOther, "/blogs/" + strconv.FormatUint(uint64(post.ID), 10) + "/update")
}




func HandlerRemoveContent(c *gin.Context) {
	postID := c.Param("postId")
	post_id, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	contentID := c.Param("contentID")
	content_id, err := strconv.ParseUint(contentID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := models.GetPostById(uint(post_id))	
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

	
	content, err := post.GetContentByID(uint(content_id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

	err = post.DeleteContent(*content)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete content"})
        return
    }

	
	c.Redirect(http.StatusSeeOther, "/blogs/" + strconv.FormatUint(uint64(post.ID), 10) + "/update")
}

