package models

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:'To Do'"`
	UserID      int64  `json:"user_id"`
}

func (u *User) AddTask(task *Task) error {
	if task == nil {
		return errors.New("task cannot be nil")
	}

	task.UserID = int64(u.ID)
	if err := db.Save(task).Error; err != nil {
		log.Printf("Error saving task for user %d: %v", u.ID, err)
		return fmt.Errorf("error saving task for user %d: %v", u.ID, err)
	}

	return nil
}

func GetUserAllTask(id int64) ([]Task, error) {
	var tasks []Task
	if err := db.Where("user_id = ?", id).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetAllTask() ([]Task, error) {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
func GetTaskById(id int64) (*Task, error) {
	var getTask Task
	if err := db.Where("ID=?", id).First(&getTask).Error; err != nil {
		return nil, err
	}
	return &getTask, nil
}

func RemoveTask(ID int64) error {
	var task Task
	if err := db.Where("ID=?", ID).Delete(&task).Error; err != nil {
		return err
	}
	return nil
}