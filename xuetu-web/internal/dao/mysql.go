package dao

import (
	"fmt"
	"log"
	"time"
	"xuetu-web/config"
	"xuetu-web/internal/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var DB *gorm.DB

func InitMySQL() {
	// 读取配置文件中的数据库连接
	cfg := config.AppConfig.MySQL
	// 构造dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect MySQL: %v", err)
	}

	// 拿到通用的数据库对象，做一些额外的数据库配置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 将其转换为全局变量，方便后续的使用
	global.Db = db
}
