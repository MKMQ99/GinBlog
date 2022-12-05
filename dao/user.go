package dao

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
)

const secret = "Monkey"

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
func CheckUpUser(id uint, name string) (code int) {
	var user model.User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == id {
		return errmsg.SUCCSE
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *model.User) int {
	data.Password = encryptPassword(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write([]byte(oPassword))
	return hex.EncodeToString(h.Sum(nil))
}

// 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]model.User, int64) {
	var users []model.User
	var total int64
	if username != "" {
		db.Select("id,username,role,created_at").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	db.Model(&users).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编辑用户信息
func EditUser(id int, data *model.User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&model.User{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user model.User
	user.ID = uint(id)
	err := db.Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 登录验证
func CheckLogin(username string, password string) (model.User, int) {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if encryptPassword(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}
	return user, errmsg.SUCCSE
}

// GetUser 查询用户
func GetUser(id int) (model.User, int) {
	var user model.User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

// ChangePassword 修改密码
func ChangePassword(id int, data *model.User) int {
	var maps = make(map[string]interface{})
	maps["password"] = data.Password

	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (model.User, int) {
	var user model.User

	db.Where("username = ?", username).First(&user)

	if encryptPassword(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCSE
}
