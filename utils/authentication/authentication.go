package authentication

import (
	"be-golang-echo/utils/config_variable"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func IsAuthenticated() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config_variable.Secret),
	})
}

func IsRefreshAuthenticated() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config_variable.RefreshSecret),
	})
}
