package main

import (
	"example.com/user-management/config"

	"example.com/user-management/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	config.SetupLogger()
	config.SetupViper()
	config.SetupDatabase()
}

func main() {
	// Initialize echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	port := viper.GetString("api.port")

	// Assign HTTP
	handler.New(e)

	// Serve server
	e.Logger.Fatal(e.Start(port))
}
