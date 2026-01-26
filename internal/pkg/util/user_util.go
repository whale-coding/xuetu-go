package util

import (
	"github.com/gin-gonic/gin"
)

// GetUserId 从Gin Context中获取用户ID
func GetUserId(c *gin.Context) (uint, bool) {
	uid, ok := c.Get("user_id")
	if !ok {
		return 0, false
	}
	userId, ok := uid.(uint)
	return userId, ok
}

// 使用demo
// GetUserInfo 获取用户信息：从Context取用户ID，无需前端传参
//func GetUserInfo(c *gin.Context) {
//	// 1. 从Context获取用户ID（封装的工具函数）
//	userId, ok := util.GetUserId(c)
//	if !ok {
//		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "未获取到用户信息，请重新登录"})
//		return
//	}
//
//	// 2. 根据用户ID查询用户信息（生产替换为GORM/MySQL查询）
//	var user *model.User
//	for _, u := range model.UserDB {
//		if u.ID == userId {
//			user = u
//			break
//		}
//	}
//	if user == nil {
//		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "用户不存在"})
//		return
//	}
//
//	// 3. 返回用户信息（密码已隐藏）
//	c.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "获取用户信息成功",
//		"data": user,
//	})
//}
