package main

import (
	"ginCms/handlers"
	"ginCms/utils"
	"github.com/spf13/viper"
)

func main() {
	// 解析配置
	utils.ParseConfig()

	// 初始化路由引擎
	r := handlers.InitEngine()
	r.Run(viper.GetString("app.addr")) // 监听并在 0.0.0.0:8080 上启动服务
}
