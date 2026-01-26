package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"xuetu-project/internal/constant"
	"xuetu-project/internal/pkg/jwt"
	"xuetu-project/internal/pkg/redis"

	"github.com/gin-gonic/gin"
)

// JwtAuth 单Token+Redis鉴权中间件（含自动续期）
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. 接口白名单：仅登录接口放行（生产可加静态资源、健康检查）
		whiteList := map[string]bool{
			"/api/auth/register": true,
			"/api/auth/login":    true,
			"/api/auth/logout":   true,
		}
		if whiteList[c.Request.URL.Path] {
			c.Next()
			return
		}

		// 2. 从Header获取Token（规范：Bearer + 空格 + Token）
		authHeader := c.GetHeader("Authorization")
		// 请求头 判断有没有Authorization字段
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请求头缺少 Authorization"})
			c.Abort()
			return
		}
		// 拆分字符串，判断格式是否正确
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token格式错误：Bearer + 空格 + Token"})
			c.Abort()
			return
		}
		tokenStr := parts[1] // 获取 Token部分

		// 3. 校验Token是否在黑名单（已退出的Token，防复用）
		if isBlack, _ := redis.SIsMember(constant.RedisKeyBlack, tokenStr); isBlack {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token已失效，请重新登录"})
			c.Abort()
			return
		}

		// 4. 解析JWT Token，获取用户ID
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token无效/已过期：" + err.Error()})
			c.Abort()
			return
		}
		userId := claims.UserId // 从 JWT 中获取用户ID

		// 5. 校验Redis中是否存在该Token（防止JWT有效但已退出/过期）
		redisKey := constant.RedisKeyToken + tokenStr
		redisUserId, err := redis.Get(redisKey) // 从Redis 获取用户ID
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token已过期/已退出"})
			c.Abort()
			return
		}

		// 6. 校验Redis中的用户ID与JWT中的是否一致（防Token伪造）
		if redisUserId != strconv.Itoa(int(userId)) {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Token无效，用户不匹配"})
			c.Abort()
			return
		}

		// 7. 核心：自动续期！刷新Redis中Token的过期时间，重置为初始2小时
		_ = redis.Expire(redisKey, time.Duration(constant.RedisTokenExpire)*time.Second)

		// 8. 将用户ID存入Gin Context，业务接口直接获取（免传参）
		c.Set("user_id", userId)

		// 9. 校验通过，放行请求
		c.Next()
	}
}
