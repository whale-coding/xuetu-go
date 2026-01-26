package redis

import (
	"context"
	"log"
	"strconv"
	"time"
	"xuetu-project/config"

	"github.com/redis/go-redis/v9"
)

// Rdb 全局Redis客户端,通过redis.Rdb.xxxx 使用
var Rdb *redis.Client
var Ctx = context.Background()

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
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}
	
	log.Println("Successfully connected to Redis")

	// 赋值全局变量
	Rdb = RedisClient
}

// Set 通用Set方法（带过期时间）
func Set(key string, value interface{}, expire time.Duration) error {
	return Rdb.Set(Ctx, key, value, expire).Err()
}

// Get 通用Get方法
func Get(key string) (string, error) {
	return Rdb.Get(Ctx, key).Result()
}

// Del 通用Del方法
func Del(key string) error {
	return Rdb.Del(Ctx, key).Err()
}

// SAdd 集合添加（黑名单）
func SAdd(key string, members ...interface{}) error {
	return Rdb.SAdd(Ctx, key, members...).Err()
}

// SIsMember 集合是否存在（校验黑名单）
func SIsMember(key string, member interface{}) (bool, error) {
	return Rdb.SIsMember(Ctx, key, member).Result()
}

// Expire 刷新过期时间
func Expire(key string, expire time.Duration) error {
	return Rdb.Expire(Ctx, key, expire).Err()
}
