package main

import (
	"fmt"

	"github.com/bastianrob/gomono/internal/notification/configs"
	notifRepo "github.com/bastianrob/gomono/internal/notification/repositories"
	notification "github.com/bastianrob/gomono/internal/notification/services"
	"github.com/bastianrob/gomono/pkg/global"
	"github.com/bastianrob/gomono/pkg/middlewares"
	"github.com/go-redis/redis/v9"
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:        configs.App.Redis.Host,
		Password:    configs.App.Redis.Pass,
		DB:          configs.App.Redis.DB,
		ReadTimeout: -1,
	})

	redisSubscriber := notifRepo.NewRedisSubscriber(redisClient)
	notifService := notification.NewNotificationService(redisSubscriber)
	notifService.Run()

	// Start server
	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%d", global.Config.Port),
		),
	)
}
