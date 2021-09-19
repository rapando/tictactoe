package utils

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func RedisConnect(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		PoolSize: 10,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		Log("ERROR", "rds", "Unable to ping because %v", err)
		return nil, err
	}
	Log("INFO", "rds", "Connected")
	return client, nil
}

func SaveToRedis(ctx context.Context, key string, data interface{}, exp int64, cache *redis.Client) error {
	Log("INFO", "rds", "Saving data to %v", key)
	expiry := time.Duration(exp) * time.Second
	if err := cache.Set(ctx, key, data, expiry).Err(); err != nil {
		Log("ERROR", "rds", "unable to save %v because %v", key, err)
		return err
	}
	Log("INFO", "rds", "Saved %v", key)
	return nil
}

func ReadFromRedis(ctx context.Context, key string, cache *redis.Client) (string, error) {
	Log("INFO", "rds", "Reading data from %v", key)
	dataStr, err := cache.Get(ctx, key).Result()
	if err != nil {
		Log("ERROR", "rds", "unable to read %v because %v", key, err)
		return "", err
	}
	Log("INFO", "rds", "found data for %v", key)
	return dataStr, nil
}
