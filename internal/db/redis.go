package db

import (
	"context"
	"fmt"
	"time"

	"github.com/MarcelloBB/ticker/internal/config"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
	enabled     = config.LoadConfigIni("redis", "enabled", false).(bool)
)

func InitRedis() {
	if !enabled {
		fmt.Println("WARN - Redis cache is disabled in configuration")
		return
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.LoadConfigIni("redis", "host", "localhost:6379").(string),
		Password: config.LoadConfigIni("redis", "password", "").(string),
		DB:       config.LoadConfigIni("redis", "db", 0).(int),
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to Redis: %v", err))
	}
}

func SetCacheValue(ctx context.Context, key string, value interface{}) error {
	if !enabled || RedisClient == nil {
		return nil
	}
	expiration := time.Duration(config.LoadConfigIni("redis", "expiration", 10).(int)) * time.Minute
	err := RedisClient.Set(Ctx, key, value, expiration).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return err
	}
	return nil
}

func GetCacheValue(ctx context.Context, key string) (string, error) {
	if !enabled || RedisClient == nil {
		return "", nil
	}
	value, err := RedisClient.Get(Ctx, key).Result()
	if err != nil {
		fmt.Println("Error getting value:", err)
		return "", err
	}
	return value, nil
}
