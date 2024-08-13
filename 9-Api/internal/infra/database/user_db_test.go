package database

import (
	"testing"

	"br.com.adrianomenezes/cursogo/9-Api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {

	user, db, _ := CreateOneUser(t)

	var userFound entity.User
	err := db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	user, _, userDB := CreateOneUser(t)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}

func CreateOneUser(t *testing.T) (*entity.User, *gorm.DB, *User) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Adriano 1", "aa@aa.com", "123adriano456")
	userDb := NewUser(db)

	err = userDb.Create(user)
	assert.Nil(t, err)
	return user, db, userDb
}
