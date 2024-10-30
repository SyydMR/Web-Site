package models


import (
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	Text string `json:"text"`
	PostID int64 `json:"post_id"`
}
