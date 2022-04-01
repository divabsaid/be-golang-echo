package jwt

import (
	"be-golang-echo/entity/user"
	"be-golang-echo/utils"
	"be-golang-echo/utils/config_variable"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(u *user.UserLoginModel) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["admin"] = false
	if u.RoleID == 1 {
		claims["admin"] = true
	}
	claims["exp"] = time.Now().Add(time.Second * 1).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config_variable.Secret))
	if err != nil {
		return nil, err
	}
	refreshtoken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshtoken.Claims.(jwt.MapClaims)
	rtClaims["id"] = u.ID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshtoken.SignedString([]byte(config_variable.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}

func GetClaims(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(config_variable.Secret), nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil
}

func GetRefreshClaims(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(config_variable.RefreshSecret), nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil
}

func GetIDfromToken(token string) (id int, err error) {
	token, err = GetTokenfromHeader(token)
	if err != nil {
		return id, err
	}
	claims, err := GetClaims(token)
	if err != nil {
		return id, err
	}
	// valid, err := verifyTokenValidity(claims, token)
	// if !valid || err != nil {
	// 	return id, err
	// }
	idClaim, _ := claims["id"].(float64)
	id = int(idClaim)
	return id, nil
}

func GetIDfromRefreshToken(token string) (id int, err error) {
	claims, err := GetRefreshClaims(token)
	if err != nil {
		return id, err
	}
	// valid, err := verifyTokenValidity(claims, token)
	// if !valid || err != nil {
	// 	return id, err
	// }
	idClaim, _ := claims["id"].(float64)
	id = int(idClaim)
	return id, nil
}

func GetTokenfromHeader(token string) (string, error) {
	if token == "" {
		return "", errors.New(utils.AUTH_REQUIRED)
	}
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]
	return token, nil
}
