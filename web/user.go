package go_manager_web

import (
	"fmt"
	db "go_manager_db"
	utils "go_manager_utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUserHandler(c *gin.Context) {
	gormDb, err := db.InitDB()
	if err != nil {
		return
	}
	users, err := db.GetAllUser(gormDb)
	// 通用响应
	utils.R(c, err, "查询角色失败", users)
}
func AddUserHandler(c *gin.Context) {

	user := db.MalUser{}
	//绑定json和结构体
	if err := c.BindJSON(&user); err != nil {
		return
	}

	gormDb, err := db.InitDB()
	if err != nil {
		return
	}

	if err := gormDb.Create(&user).Error; err != nil {
		fmt.Println("插入数据失败")
		return
	}
	// 通用响应
	utils.R(c, err, "添加角色失败", "添加角色成功")
}
func DelUserHandler(c *gin.Context) {
	// 从url获取参数
	idStr := c.Query("uid")
	// fmt.Println(idStr)
	uid, err := strconv.ParseInt(idStr, 10, 64)
	gormDb, err := db.InitDB()
	result := gormDb.Delete(&db.MalUser{}, uid)
	if result.Error != nil {
		fmt.Println("删除数据失败")
		return
	}
	// 通用响应
	utils.R(c, err, "删除角色失败", "删除角色成功")

}
func GetOneUserHandler(c *gin.Context) {
	// 从url获取参数
	idStr := c.Query("uid")
	fmt.Println(idStr)
	uid, _ := strconv.ParseInt(idStr, 10, 64)
	gormDb, err := db.InitDB()
	if err != nil {
		return
	}
	one, err2 := db.GetUserById(gormDb, uid)
	// 通用响应
	utils.R(c, err2, "查询角色失败", one)
}
func UptUserHandler(c *gin.Context) {
	// 从url获取参数

	user := db.MalUser{}
	//绑定json和结构体
	if err := c.BindJSON(&user); err != nil {
		return
	}

	gormDb, err := db.InitDB()
	if err != nil {
		return
	}
	err = db.UptUserById(gormDb, user.ID, user)
	if err != nil {
		return
	}
	// 通用响应
	utils.R(c, err, "修改角色失败", "修改角色成功")
}
func registerUser(middles ...gin.HandlerFunc) {
	// 创建路由组v1/user
	user := DefineRouteGroup(v1, "user", r)
	// 添加中间件
	if middles != nil {
		user.Use(middles...)
	}
	user.GET("all", GetAllUserHandler)
	// 添加
	user.POST("add", AddUserHandler)
	// 删除
	user.DELETE("del", DelUserHandler)
	// 根据id获取
	user.GET("id", GetOneUserHandler)
	// 根据id修改
	user.PUT("upt", UptUserHandler)
}
