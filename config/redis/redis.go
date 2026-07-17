package redis

import (
	"context"
	"fmt"
	"strconv"
	"ticketing-system/config"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

var Client *goredis.Client

func ConnectRedis() {
	db, err := strconv.Atoi(config.GetEnvOrDefault("REDIS_DB", "0"))
	if err != nil {
		panic("Invalid REDIS_DB")
	}

	Client = goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.GetEnvOrPanic("REDIS_HOST"), config.GetEnvOrPanic("REDIS_PORT")),
		Password: config.GetEnvOrDefault("REDIS_PASSWORD", ""),
		DB:       db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		panic(err)
	}
}

func CloseRedis() {
	if Client != nil {
		_ = Client.Close()
	}
}
