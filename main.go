package main

import (
	"be-golang-echo/app"
	"be-golang-echo/utils/config_variable"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	//init config
	app.InitConfig()
	config_variable.Secret = os.Getenv("JWT_SECRET")
	config_variable.RefreshSecret = os.Getenv("JWT_REFRESH_SECRET")

	// init database
	dbConn := app.InitDatabase()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
	}))

	// init entity
	app.InitEntity(e, dbConn)

	// Init logger
	zerolog.TimestampFieldName = "timestamp"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
