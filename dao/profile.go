package dao

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
)

// GetProfile 获取个人信息设置
func GetProfile(id int) (model.Profile, int) {
	var profile model.Profile
	err = db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCSE
}

// UpdateProfile 更新个人信息设置
func UpdateProfile(id int, data *model.Profile) int {
	var profile model.Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
