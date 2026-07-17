package database

import (
	"fmt"
	"ticketing-system/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetEnvOrPanic("DB_HOST"),
		config.GetEnvOrPanic("DB_USERNAME"),
		config.GetEnvOrPanic("DB_PASSWORD"),
		config.GetEnvOrPanic("DB_NAME"),
		config.GetEnvOrPanic("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
}
