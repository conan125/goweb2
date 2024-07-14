package go_manager_db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

// 查数据

func GetAllRole(db *gorm.DB) ([]*MalRole, error) {
	var roles []*MalRole
	err := db.Find(&roles).Error
	if err != nil {
		fmt.Println("查询信息失败")
		fmt.Println(err)
		return nil, err
	}
	return roles, nil
}

// 根据id查数据

func GetRoleById(db *gorm.DB, id int64) (*MalRole, error) {
	var role MalRole
	err := db.First(&role, id).Error
	if err != nil {
		fmt.Println("查询信息失败")
		return nil, err
	}
	return &role, nil
}

// 根据id改数据

func UptRoleById(db *gorm.DB, id int64, roleName string) error {
	err := db.Model(&MalRole{}).Where("id = ?", id).Update("role", roleName).Error
	if err != nil {
		fmt.Println("修改信息失败")
		return err
	}
	return nil
}
