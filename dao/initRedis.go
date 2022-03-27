package dao

import "github.com/go-redis/redis"

var redisDB *redis.Client

func InitRedis() error {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
