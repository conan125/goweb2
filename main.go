package main

import (
	db "go_manager_db"
	web "go_manager_web"
)

func main() {

	// 应用程序退出时关闭数据库连接
	// 初始化数据库
	gormDb, err := db.InitDB()
	if err != nil {
		panic("failed to migrate database")
	}
	err = gormDb.AutoMigrate(&db.MalUser{})
	if err != nil {
		panic("failed to migrate database")
	}

	// 开启服务
	web.Run()
}
