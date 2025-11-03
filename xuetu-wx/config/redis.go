package config

import (
	"log"
	"xuetu-wx/global"

	"github.com/go-redis/redis"
)

// RedisDB 全局变量定义
//var RedisDB *redis.Client
//var (
//	rdb *redis.Client
//	ctx = context.Background()
//)

// InitRedis 初始化Redis
func InitRedis() {
	addr := AppConfig.Redis.Addr
	db := AppConfig.Redis.DB
	password := AppConfig.Redis.Password

	// 关键点：初始化redisClient
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})

	// 测试redis是否能够连通
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	// 绑定为全局变量，方便后续使用
	global.RedisDB = RedisClient
}
