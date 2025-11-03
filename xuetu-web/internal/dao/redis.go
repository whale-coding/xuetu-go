package dao

import (
	"log"
	"xuetu-web/config"
	"xuetu-web/internal/global"

	"github.com/go-redis/redis"
)

// RedisDB 全局变量定义
//var RDB *redis.Client
//var ctx = context.Background()

// InitRedis 初始化Redis
func InitRedis() {
	cfg := config.AppConfig.Redis

	// 关键点：初始化redisClient
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试redis是否能够连通
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	// 绑定为全局变量，方便后续使用
	global.RedisDB = RedisClient
}
