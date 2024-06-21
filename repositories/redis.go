package repositories

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func CreateConnectionRedis() *redis.Client {
	redis_db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		DB:   redis_db,
	})
}
