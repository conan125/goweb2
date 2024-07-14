package go_manager_db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

// 查数据
func GetAllUser(db *gorm.DB) (users []MalUser, err error) {
	err = db.Find(&users).Error
	if err != nil {
		fmt.Println("查询信息失败")
		return nil, err
	}
	return users, nil
}

// 根据ID查数据
func GetUserById(db *gorm.DB, id int64) (user MalUser, err error) {
	err = db.First(&user, id).Error
	if err != nil {
		fmt.Println("查询信息失败")
		return MalUser{}, err
	}
	return user, nil
}

// 根据Name和Pass查数据
func GetUserByName(db *gorm.DB, uname string, upass string) (user MalUser, err error) {
	err = db.Where("uname = ? AND upass = ?", uname, upass).First(&user).Error
	if err != nil {
		fmt.Println("查询信息失败")
		return MalUser{}, err
	}
	return user, nil
}

// 根据ID更新数据
func UptUserById(db *gorm.DB, id int64, updates interface{}) (err error) {
	err = db.Model(&MalUser{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		fmt.Println("修改信息失败")
		return err
	}
	return nil
}
