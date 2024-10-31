package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	PostID      uint  `json:"post_id"`
	Type        string `json:"type"`
	Data        string `json:"context"`
}



func GetAllContentByPostID(postID uint) ([]Content, error) {
	var contents []Content

	if err := db.Where("post_id = ?", postID).Order("id ASC").Find(&contents).Error; err != nil {
		return nil, fmt.Errorf("error fetching contents for post ID %d: %w", postID, err)
	}

	return contents, nil
}


func (p *Post) CreateContent(content Content) error {
    content.PostID = p.ID

    if err := db.Create(&content).Error; err != nil {
        return fmt.Errorf("error creating content: %w", err)
    }

    p.Contents = append(p.Contents, content)
    return nil
}


func (p *Post) DeleteContent(content Content) error {
    if err := db.Delete(&content).Error; err != nil {
        return fmt.Errorf("error deleting content: %w", err)
    }
    for i, c := range p.Contents {
        if c.ID == content.ID {
            p.Contents = append(p.Contents[:i], p.Contents[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("content with ID %d not found in post contents", content.ID)
}



func (p *Post) GetContentByID(id uint) (*Content,  error) {
	var content Content

	if err := db.Where("ID=?", id).First(&content).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, fmt.Errorf("content with ID %d not found: %w", id, err)
        }
        return nil, fmt.Errorf("error fetching content for content ID %d: %w", id, err)	
	}

	return &content, nil
}
