package main

import (
	"be-golang-echo/app"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	//init config
	config := app.InitConfig()

	// init database
	dbConn := app.InitDatabase(config)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
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
	e.Logger.Fatal(e.Start(":" + config.GetString("port")))
}
