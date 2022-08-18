package main

import (
	"fmt"

	"github.com/bastianrob/gomono/internal/credential/configs"
	credential "github.com/bastianrob/gomono/internal/credential/controllers"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/bastianrob/gomono/pkg/middlewares"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	godotenv.Load()
	global.Init()
	configs.Init()
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		ContentSecurityPolicy: "default-src 'self'",
		HSTSMaxAge:            3600,
	}))
	e.Use(middleware.BodyLimit("100KB"))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middlewares.RequestLogger())

	credController := credential.InitializeController()
	credController.Routes(e)

	// Start server
	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%d", global.Config.Port),
		),
	)
}
