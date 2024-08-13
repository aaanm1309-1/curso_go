package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Adriano", "a@a.com.br", "123mamae456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Name)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Password)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Adriano", user.Name)
	assert.Equal(t, "a@a.com.br", user.Email)
	assert.NotEmpty(t, user.ID)
	assert.True(t, user.ValidatePassword("123mamae456"))
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Adriano", "a@a.com.br", "123mamae456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123mamae456"))
	assert.False(t, user.ValidatePassword("123mamae4567"))
	assert.NotEqual(t, user.Password, ("123mamae456"))

}
