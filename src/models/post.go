package models


import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Author User `json:"author"`
	Tags []Tags `json:"tags"`
	Contents []Content `json:"contents"`
}


func GetAllPost() ([]Post, error) {
	var posts []Post
	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostById(id int64) (*Post, error) {
	var getPost Post
	if err := db.Where("ID=?", id).First(&getPost).Error; err != nil {
		return nil, err
	}
	return &getPost, nil
}