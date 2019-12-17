package app

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

//初始化redis
func initRedis(){

	redisConfig := Cfg.Section("redis")
	// 典型读取操作，默认分区可以使用空字符串表示
	host := redisConfig.Key("host").String()
	port := redisConfig.Key("port").String()

	redisAddr := fmt.Sprintf("%v:%v", host, port)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
