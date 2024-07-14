package go_manager_web

import (
	"fmt"
	db "go_manager_db"
	utils "go_manager_utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllRoleHandler(c *gin.Context) {
	gormDb, err := db.InitDB()
	roles, err := db.GetAllRole(gormDb)
	// 通用响应
	utils.R(c, err, "获取角色列表失败", roles)
}
func AddRoleHandler(c *gin.Context) {
	//绑定json和结构体
	var gormDb *gorm.DB
	role := db.MalRole{}
	//绑定json和结构体
	if err := c.BindJSON(&role); err != nil {
		return
	}
	gormDb, err := db.InitDB()
	if err != nil {
		return
	}

	if err := gormDb.Create(&role).Error; err != nil {
		fmt.Println("插入数据失败")
		return
	}
	// 通用响应
	utils.R(c, err, "添加角色失败", "添加角色成功")

}
func DelRoleHandler(c *gin.Context) {
	// 从url获取参数
	idStr := c.Query("rid")
	// fmt.Println(idStr)
	rid, err := strconv.ParseInt(idStr, 10, 64)
	gormDb, err := db.InitDB()
	if err != nil {
		return
	}
	gormDb.Delete("mal_role", rid)
	// 通用响应
	utils.R(c, err, "删除角色失败", "删除角色成功")
}
func GetOneRoleHandler(c *gin.Context) {
	// 从url获取参数
	idStr := c.Query("rid")
	fmt.Println(idStr)
	rid, _ := strconv.ParseInt(idStr, 10, 64)
	gormDb, err := db.InitDB()
	if err != nil {
		return
	}
	one, err2 := db.GetRoleById(gormDb, rid)
	// 通用响应
	utils.R(c, err2, "查询角色失败", one)
}
func UptRoleHandler(c *gin.Context) {
	role := db.MalRole{}
	//绑定json和结构体
	if err := c.BindJSON(&role); err != nil {
		return
	}
	gormDb, err := db.InitDB()
	if err != nil {
		return
	}
	err = db.UptRoleById(gormDb, role.ID, role.Role)
	// 通用响应
	utils.R(c, err, "修改角色失败", "修改角色成功")
}
func registerRole(middles ...gin.HandlerFunc) {
	// 创建路由组v1/user
	role := DefineRouteGroup(v1, "role", r)
	// 添加中间件
	if middles != nil {
		role.Use(middles...)
	}
	// 获取所有
	role.GET("all", GetAllRoleHandler)
	// 添加
	role.POST("add", AddRoleHandler)
	// 删除
	role.DELETE("del", DelRoleHandler)
	// 根据id获取
	role.GET("id", GetOneRoleHandler)
	// 根据id修改
	role.PUT("upt", UptRoleHandler)
}
