package main

import (
	"fmt"
	"xuetu-wx/config"
	"xuetu-wx/router"
)

func main() {
	// 配置信息
	config.InitConfig()

	port := config.AppConfig.App.Port
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server started on port %s\n", port)

	// 路由信息
	r := router.SetupRouter()

	// 绑定端口，运行
	r.Run("127.0.0.1:" + port)
}
