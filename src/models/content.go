package models

import (
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	PostID      int64  `json:"post_id"`
	Type        string `json:"type"`
	Data        string `json:"context"`
	OrderNumber int64  `json:"order_number"`
}



