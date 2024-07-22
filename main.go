package main

import (
	"ginCms/handlers/system"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", system.Ping)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
