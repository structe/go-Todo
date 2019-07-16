package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email",gorm:"unique_index;default:null`
	Password string `json:"password",gorm:"default:null`
}

// User
func (user *User) Insert() error {
	return DB.Create(user).Error
}

func (user *User) Update() error {
	return DB.Save(user).Error
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.First(&user, "email = ?", email).Error
	return &user, err
}

func GetUser(id interface{}) (*User, error) {
	var user User
	err := DB.First(&user, id).Error
	return &user, err
}
