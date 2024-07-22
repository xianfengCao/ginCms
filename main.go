package main

import (
	"ginCms/handlers"
)

func main() {
	// 初始化路由引擎
	r := handlers.InitEngine()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
