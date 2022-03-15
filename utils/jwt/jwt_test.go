package jwt_test

import (
	"be-golang-echo/entity/user"
	"be-golang-echo/utils/config_variable"
	"be-golang-echo/utils/jwt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userLogin = &user.UserLoginModel{
		ID:       1,
		Username: "User1",
		Password: "password",
		RoleID:   1,
	}

	wrongToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6NDc1Njg3NTAwOSwiaWQiOjR9.bJBQq5NeG14xwWOVFBum0w8RfpTjBXQYtM--bnYWflE"
	token      = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjo0NzU2ODIzMzEzLCJpZCI6Mn0.Bzz5j-okD6D7obfYomt03kHVmvl4nUB0-ROEEQU1TGA"

	tokenHeader      = "Bearer " + token
	wrongTokenHeader = "Bearer " + wrongToken
)

func TestCreateJWT(t *testing.T) {
	token, err := jwt.CreateJWTToken(userLogin)
	assert.NotNil(t, token)
	assert.NoError(t, err)

}

func TestGetClaimsError(t *testing.T) {
	config_variable.Secret = "error"
	_, err := jwt.GetClaims(token)
	assert.Error(t, err)
}

func TestVerifyAdminToken(t *testing.T) {
	config_variable.Secret = ""
	// r := redis.InitRedisTest()
	// r.On("Get", "\x02").Return(redisClient.NewStringResult(token, nil))
	// redis.Rdb = r
	admin, err := jwt.VerifyAdminToken(tokenHeader)
	assert.True(t, admin)
	assert.NoError(t, err)

}

func TestVerifyAdminTokenEmpty(t *testing.T) {
	admin, err := jwt.VerifyAdminToken("")
	assert.False(t, admin)
	assert.Error(t, err)

}

func TestVerifyAdminTokenErrSign(t *testing.T) {
	config_variable.Secret = "error"
	admin, err := jwt.VerifyAdminToken(tokenHeader)
	assert.False(t, admin)
	assert.Error(t, err)

}

func TestVerifyAdminTokenNotAdmin(t *testing.T) {
	config_variable.Secret = ""
	// r := redis.InitRedisTest()
	// r.On("Get", "\x04").Return(redisClient.NewStringResult(wrongToken, nil))
	// redis.Rdb = r
	admin, err := jwt.VerifyAdminToken(wrongTokenHeader)
	assert.False(t, admin)
	assert.Error(t, err)

}


func TestGetIDfromToken(t *testing.T) {
	// config_variable.Secret = ""
	// r := redis.InitRedisTest()
	// r.On("Get", "\x02").Return(redisClient.NewStringResult(token, nil))
	// redis.Rdb = r
	id, err := jwt.GetIDfromToken(tokenHeader)
	assert.Equal(t, 2, id)
	assert.NoError(t, err)

}

func TestGetIDfromTokenEmpty(t *testing.T) {
	_, err := jwt.GetIDfromToken("")
	assert.Error(t, err)

}

func TestGetIDfromTokenErrSign(t *testing.T) {
	config_variable.Secret = "error"
	_, err := jwt.GetIDfromToken(tokenHeader)
	assert.Error(t, err)
}
