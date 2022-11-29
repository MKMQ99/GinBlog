package dao

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
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

// 新增用户
func CreateUser(data *model.User) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
