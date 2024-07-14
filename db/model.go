package go_manager_db

import "gorm.io/gorm"

// 专门定义与数据库交互的结构体

// 用户表
type MalUser struct {
	gorm.Model
	ID    int64  `gorm:"column:id;primaryKey;autoIncrement" json:"Id"`
	Uname string `gorm:"column:uname" json:"Uname"`
	Upass string `gorm:"column:upass" json:"Upass"`
	Rid   int64  `gorm:"column:rid" json:"Rid"`
}

// 角色表

type MalRole struct {
	gorm.Model
	ID   int64  `gorm:"primaryKey;autoIncrement"`
	Role string `gorm:"type:varchar(20);not null;unique;default:user"`
}
