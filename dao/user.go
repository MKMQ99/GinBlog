package dao

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 查询用户是否存在
func CheckUser(username string) int {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user model.User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCSE
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *model.User) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) []model.User {
	var users []model.User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户
func EditUser(c *gin.Context) {
	
}
