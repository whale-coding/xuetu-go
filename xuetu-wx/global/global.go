package global

import (
	"github.com/go-redis/redis"
)

// 全局变量定义
var (
	RedisDB *redis.Client
)
