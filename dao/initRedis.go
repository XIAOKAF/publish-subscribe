package dao

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

var redisDB *redis.Client

func InitRedis() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     15, //连接池最大连接数
		MinIdleConns: 5,  //闲置连接数量
		//超时设置
		DialTimeout:  5 * time.Second, //连接建立的超时时间
		ReadTimeout:  3 * time.Second, //读超时，-1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时
		PoolTimeout:  4 * time.Second, //党所有连接都处在繁忙的状态时，客户端等待可用连接的最大时间
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	defer redisDB.Close()
}
