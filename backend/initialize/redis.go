package initialize

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func initRedis() {
	ctx := context.Background()
	db := viper.GetInt("Redis.db")
	host := viper.GetString("Redis.host")
	port := viper.GetInt("Redis.port")
	password := viper.GetString("Redis.password")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,
	})
	result, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	slog.Info("redis 连接成功", "result", result)
	global.RedisClient = client
}
