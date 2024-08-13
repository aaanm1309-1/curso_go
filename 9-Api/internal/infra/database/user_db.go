package database

import (
	"br.com.adrianomenezes/cursogo/9-Api/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {

	return &User{DB: db}
}

func (user *User) Create(newUser *entity.User) error {
	return user.DB.Create(newUser).Error
}

func (user *User) FindByEmail(email string) (*entity.User, error) {
	var userReturned entity.User
	if err := user.DB.Where("email = ?", email).First(&userReturned).Error; err != nil {
		return nil, err
	}
	return &userReturned, nil
}
