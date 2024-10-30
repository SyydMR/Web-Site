package models


import (
	"gorm.io/gorm"
)


type Content struct {
	gorm.Model
	PostID int64 `json:"post_id"`
	Type string `json:"type"`
	Context string `json:"context"`
	OrderNumber int64 `json:"order_number"`
}