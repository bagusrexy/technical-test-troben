package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitRedis(addr, password string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func SetCache(key string, value []byte, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

func GetCache(key string) ([]byte, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key does not exist")
	} else if err != nil {
		return nil, err
	}
	return []byte(val), nil
}
