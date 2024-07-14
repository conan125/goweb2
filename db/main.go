package go_manager_db

import (
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据库相关操作

// 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	dsn := "./manager.db"
	gormDb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// 连接
	// Open可能仅校验参数，而没有与db间创建连接，
	// 要确认db是否可用，需要调用Ping。Connect则相当于Open+Ping。
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return nil, err
	}
	CreateRoleTable(gormDb)
	// 设置连接池
	sqlDB, err := gormDb.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(100)

	// 设置连接可复用的最长时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 测试数据库连接
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("database connection successfully configured")
	return gormDb, err
}

// 创建用户表
type Role struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Role string `gorm:"type:varchar(20);not null;unique;default:user"`
}

func CreateRoleTable(godb *gorm.DB) error {
	// 自动迁移模式，创建表
	err := godb.AutoMigrate(&MalRole{})
	if err != nil {
		return err
	}

	// 初始化表数据
	roles := []Role{
		{Role: "user"},
		{Role: "admin"},
		{Role: "super"},
		{Role: "root"},
	}
	for _, role := range roles {
		var existingRole MalRole
		result := godb.Where(&MalRole{Role: role.Role}).FirstOrCreate(&existingRole, role)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
