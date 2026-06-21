package config

import "os"

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Missing environment variable: " + key)
	}
	return value
}
