package router

import (
	"time"
	"xuetu-wx/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 初始化Gin
	r := gin.Default()

	// 跨域配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 路由分组  /api/auth 前缀提取出来公共的前缀
	auth := r.Group("/api/wx")
	{
		auth.GET("/callback", controller.VerifySignature)
		auth.POST("/callback", controller.CallbackHandler)
	}

	// 返回r
	return r
}
