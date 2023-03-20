package connections

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func NewRedis() (redis.UniversalClient, error) {
	redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{redisAddr},
		Password: "",
		DB:       0,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return client, nil
}
