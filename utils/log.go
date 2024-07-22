package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"log"
	"log/slog"
	"os"
	"time"
)

// SetLogger 设置日志
func SetLogger() {
	// 1. 设置集中的日志Writer
	setLoggerWriter()

	// 2. 初始化
	initLogger()
}

// logger
var logger *slog.Logger

func Logger() *slog.Logger {
	return logger
}

// 公共的writer变量
var logWriter io.Writer

func LogWriter() io.Writer {
	return logWriter
}

// 设置writer
func setLoggerWriter() {
	// 根据不同的mode，选择不同的writer
	switch gin.Mode() {
	case gin.ReleaseMode:
		// 打开文件
		month := time.Now().Format("200601")
		logfile := viper.GetString("app.log.path")
		logfile += fmt.Sprintf("/app-%s.log", month)

		if file, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
			log.Println(err)
			return
		} else {
			logWriter = file
		}
	case gin.TestMode, gin.DebugMode:
		fallthrough
	default:
		logWriter = os.Stdout
	}
}

// 初始化日志
func initLogger() {
	// 使用json模式记录
	logger = slog.New(slog.NewJSONHandler(logWriter, &slog.HandlerOptions{}))
}
