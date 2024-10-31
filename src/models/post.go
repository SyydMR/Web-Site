package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	AuthorID   uint      `json:"author_id"`
	// Tags     []Tag     `json:"tags" gorm:"many2many:post_tags"`
	Contents []Content `json:"contents" gorm:"foreignKey:PostID"`
	Publish  bool      `json:"publish"`
}

func GetAllPost() ([]Post, error) {
	var posts []Post
	if err := db.Preload("Contents").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostById(id uint) (*Post, error) {
	var getPost Post
	if err := db.Preload("Contents").Where("ID=?", id).First(&getPost).Error; err != nil {
		return nil, err
	}
	return &getPost, nil
}



func CreatePost(post *Post) (error) {
	if err := db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}


func GetPostsByUserID(userID uint) ([]Post, error) {
	var posts []Post

	if err := db.Preload("Contents").Where("author_id = ?", userID).Find(&posts).Error; err != nil {
        return nil, fmt.Errorf("error fetching posts for user ID %d: %w", userID, err)
	}

	return posts, nil
}



func DeletePost(postID uint) error {
	if err := db.Where("ID = ?", postID).Delete(&Post{}).Error; err != nil {
		return fmt.Errorf("error deleting post with ID %d: %w", postID, err)
	}
	return nil
}