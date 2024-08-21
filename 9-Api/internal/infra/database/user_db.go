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

func (user *User) FindAll(page, limit int, sort string) ([]entity.User, error) {
	var usersReturned []entity.User
	var err error
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && page > 0 && limit != 0 && limit > 0 {
		err = user.DB.Limit(limit).Offset((page - 1) * limit).Order("name " + sort).Find(&usersReturned).Error
	} else {
		err = user.DB.Order("name " + sort).Find(&usersReturned).Error
	}
	return usersReturned, err
}
