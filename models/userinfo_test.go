package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	id       = "1"
	email    = "test@test.com"
	username = "test_username"
	fullName = "Test Name"
	photoURL = "test.png"
)

func TestCreateUserinfo(t *testing.T) {
	userinfo := CreateUserinfo(id, email, username, fullName, photoURL)

	assert := assert.New(t)
	assert.NotNil(userinfo)
}
