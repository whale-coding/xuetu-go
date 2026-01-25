package router

import (
	"time"
	"xuetu-project/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(ctrl *controller.Controller) *gin.Engine {
	// 初始化 gin引擎
	r := gin.Default()

	// 跨域配置需要放在路由前面，否则不生效！
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API 版本
	api := r.Group("/api")

	// 认证相关路由 (公开)
	auth := api.Group("/auth")
	{
		auth.POST("/login", ctrl.UserController.Login)
		auth.POST("/register", ctrl.UserController.Register)
	}

	// 用户管理路由
	user := api.Group("/user")
	{
		user.POST("", ctrl.UserController.CreateUser)                            // 创建用户
		user.GET("/:id", ctrl.UserController.GetUserByID)                        // 根据ID 获取用户
		user.GET("", ctrl.UserController.QueryUsers)                             // 分页查询用户列表
		user.POST("/querywithcond", ctrl.UserController.QueryUsersWithCondition) // 条件查询用户列表
		user.GET("/search", ctrl.UserController.GetUserByUsername)               // 根据用户名获取用户
		user.PUT("/:id", ctrl.UserController.UpdateUser)                         // 更新用户
		user.DELETE("/:id", ctrl.UserController.DeleteUser)                      // 删除用户
		user.PUT("/:id/status", ctrl.UserController.UpdateUserStatus)            // 更新用户状态
		user.POST("/batch-delete", ctrl.UserController.BatchDeleteUsers)         // 批量删除用户
	}

	// 返回 gin引擎
	return r
}
