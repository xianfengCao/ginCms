package models

import (
	"ginCms/utils"
	"log"
)

// Init 初始化模型
func Init() {
	// migrate
	migrate()

	// seed
	seed()
}

// 表结构迁移
func migrate() {
	// 自动迁移
	if err := utils.DB().AutoMigrate(
		&Role{},
	); err != nil {
		log.Fatalln(err)
	}
}

// 数据填充
func seed() {
	roleSeed()
}
