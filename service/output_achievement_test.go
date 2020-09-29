package service

import (
	"testing"

	"local.packages/models"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	user := models.User
	user.Email = "test@mail.com"
	user.Password = "password0021"

	userService := &UserService{}

	err := userService.AddTodos(user[:])

	assert.Nil(t, err)
}
