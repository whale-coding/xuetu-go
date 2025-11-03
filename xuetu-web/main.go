package main

import (
	"fmt"
	"xuetu-web/config"
	"xuetu-web/internal/dao"
	"xuetu-web/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode("release") // 设置简约日志打印

	config.InitConfig() // 读取配置文件

	port := config.AppConfig.Server.Port
	fmt.Printf("Server started on port %s\n", port)

	dao.InitMySQL() // 初始化 MySQL
	dao.InitRedis() // 初始化 Redis

	// 路由信息
	r := router.SetupRouter()

	// 绑定端口,运行
	r.Run(":8080")
}

//go get github.com/gin-gonic/gin
//go get github.com/spf13/viper
//go get github.com/go-redis/redis
//go get gorm.io/gorm
//go get gorm.io/driver/mysql
//go get github.com/gin-contrib/cors
