package password_test

import (
	"be-golang-echo/utils/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

var hashPwd = "$2a$04$hfo67dm4.grZgGblhRh/9.C9RN6DM43EsQiWdYXjxbkfAHMR7xLNy"

func TestHashPassword(t *testing.T) {
	hash, err := password.HashPassword("pwd123")
	assert.NotNil(t, hash)
	assert.NoError(t, err)

}

func TestVerifyPassword(t *testing.T) {
	hash, err := password.VerifyPassword("admin", hashPwd)
	assert.NotNil(t, hash)
	assert.NoError(t, err)

}

func TestVerifyPasswordNotMatch(t *testing.T) {
	_, err := password.VerifyPassword("admi", hashPwd)
	assert.Error(t, err)

}
