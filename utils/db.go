package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// InitDB 初始化数据库连接
func InitDB() {
	// 1. 配置
	logLevel := gormLogger.Warn
	// 根据应用的mod，控制级别
	switch gin.Mode() {
	case gin.ReleaseMode:
		logLevel = gormLogger.Warn
	case gin.TestMode, gin.DebugMode:
		fallthrough
	default:
		// 最多的日志
		logLevel = gormLogger.Info
	}
	// db 日志
	unionLogger := gormLogger.New(
		log.New(LogWriter(), "\n", log.LstdFlags),
		gormLogger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logLevel,
		},
	)
	// gorm 连接配置
	conf := &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
		}, // 数据表命名策略
		FullSaveAssociations:                     false,
		Logger:                                   unionLogger,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true, // 数据表迁移时禁用外键约束
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		TranslateError:                           false,
		PropagateUnscoped:                        false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}

	// 2. 创建db对象
	dsn := viper.GetString("db.dsn")
	if dbNew, err := gorm.Open(mysql.Open(dsn), conf); err != nil {
		log.Fatalln(err)
	} else {
		db = dbNew
	}
}

// 全局的db对象
var db *gorm.DB

// DB 全局访问db对象的方法
func DB() *gorm.DB {
	return db
}
