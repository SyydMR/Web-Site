package models

import (
	"errors"
	"fmt"

	"github.com/SyydMR/Web-Site/src/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"default:'New User'"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique;default:null"`
	Tasks    []Task `gorm:"foreignKey:UserID" json:"tasks"`
	Posts    []Post `gorm:"foreignKey:AuthorID" json:"posts"`
}

func GetAllUser() ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id uint) (*User, error) {
	var getUser User
	if err := db.Preload("Tasks").Where("ID = ?", id).First(&getUser).Error; err != nil {
		return nil, err
	}
	return &getUser, nil
}

func (u *User) Register() (*User, error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	u.Password = hashedPassword

	result := db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (u *User) Login(plainPassword string) (string, error) {
	var user User

	result := db.Where("username = ?", u.Username).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	if !utils.CheckPasswordHash(plainPassword, user.Password) {
		return "", errors.New("incorrect password")
	}
	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func DeleteUser(int uint) error {
	return db.Delete(&User{Model: gorm.Model{ID: int}}).Error
}

func DeleteAllUsers() error {
    if err := db.Delete(&User{}, "1 = 1").Error; err != nil {
        return fmt.Errorf("failed to delete users: %w", err)
    }
    return nil
}