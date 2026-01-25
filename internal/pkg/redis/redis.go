package redis

import (
	"log"
	"strconv"
	"xuetu-project/config"

	"github.com/go-redis/redis"
)

// RDB 全局Redis客户端,通过redis.RDB.xxxx 使用
var RDB *redis.Client

// InitRedis 初始化Redis
func InitRedis() {
	cfg := config.AppConfig.Redis
	
	// 关键点：初始化redisClient
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试 redis 是否能够连通
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	log.Println("Successfully connected to Redis")

	// 赋值全局变量
	RDB = RedisClient
}
