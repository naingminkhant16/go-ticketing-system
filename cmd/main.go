package main

import (
	"os"
	"strconv"
	"ticketing-system/config"
	"ticketing-system/config/cloud"
	"ticketing-system/config/database"
	"ticketing-system/config/database/migration"
	"ticketing-system/config/database/seeder"
	redisconfig "ticketing-system/config/redis"
	"ticketing-system/handler"
	"ticketing-system/middleware"
	"ticketing-system/repository"
	"ticketing-system/route"
	"ticketing-system/service"
	"ticketing-system/service/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load local env values when present
	_ = godotenv.Load()

	router := gin.Default()

	// middlewares
	router.Use(middleware.ErrorHandler())

	// connect database
	database.ConnectDatabase()

	// connect redis
	redisconfig.ConnectRedis()
	defer redisconfig.CloseRedis()

	// database migration
	migration.MigrateDatabase(database.DB)

	// seed admin users
	_ = seeder.SeedAdminUsers(database.DB)

	// register module
	registerModules(router)

	// load S3
	cloud.LoadS3Config()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		panic(err)
		return
	}
}

func registerModules(router *gin.Engine) {
	smtpPort, err := strconv.Atoi(config.GetEnvOrPanic("SMTP_PORT"))
	if err != nil {
		panic("Invalid SMTP_PORT")
	}
	// mail service
	mailService := service.NewSMTPService(config.SMTPConfig{
		Host:     config.GetEnvOrPanic("SMTP_HOST"),
		Port:     smtpPort,
		Username: config.GetEnvOrPanic("SMTP_USERNAME"),
		Password: config.GetEnvOrPanic("SMTP_PASSWORD"),
		From:     config.GetEnvOrPanic("SMTP_FROM_MAIL_ADDRESS"),
	})

	// s3 service
	s3 := storage.NewS3(cloud.S3Client)

	// users
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// auth
	authService := service.NewAuthService(userService, mailService)
	authHandler := handler.NewAuthHandler(authService)

	// events
	eventRepository := repository.NewEventRepository(database.DB)
	eventService := service.NewEventService(eventRepository, s3)
	eventHandler := handler.NewEventHandler(eventService)

	route.RegisterAuthRoutes(router, authHandler)
	route.RegisterUserRoutes(router, userHandler)
	route.RegisterEventRoutes(router, eventHandler)
}
