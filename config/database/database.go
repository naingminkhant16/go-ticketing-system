package database

import (
	"fmt"
	"ticketing-system/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnvOrPanic("DB_USERNAME"),
		config.GetEnvOrPanic("DB_PASSWORD"),
		config.GetEnvOrPanic("DB_HOST"),
		config.GetEnvOrPanic("DB_PORT"),
		config.GetEnvOrPanic("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database")
	}
	
	DB = db
}
