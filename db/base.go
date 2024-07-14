package go_manager_db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

// 插入数据
func InsertUser(db *gorm.DB, user *MalUser) error {
	if err := db.Create(user).Error; err != nil {
		fmt.Println("插入数据失败")
		return err
	}
	return nil
}

// 删除数据
func DeleteUser(db *gorm.DB, id int64) error {
	if err := db.Delete(&MalUser{}, id).Error; err != nil {
		fmt.Println("删除数据失败")
		return err
	}
	return nil
}
