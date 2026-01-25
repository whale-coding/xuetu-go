package main

import (
	"fmt"
	"xuetu-project/config"
	"xuetu-project/internal/controller"
	"xuetu-project/internal/pkg/mysql"
	"xuetu-project/internal/pkg/redis"
	"xuetu-project/internal/repository"
	"xuetu-project/internal/router"
	"xuetu-project/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("release") // 设置 gin为发布模式

	config.InitConfig() // 初始化配置
	// 读取端口号
	port := config.AppConfig.Server.Port
	fmt.Printf("Server started on port %s\n", port)

	// 初始化数据库 - 返回数据库实例
	db := mysql.InitMySQL()

	// 初始化 Redis
	redis.InitRedis()

	// 初始化 RabbitMQ
	//rabbitmq.InitRabbitMQ()
	//defer rabbitmq.Close() // 确保程序退出时关闭连接

	// 初始化依赖
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc)

	// 路由信息
	r := router.SetupRouter(ctrl)

	// 绑定端口,运行
	r.Run(":" + port)
}
