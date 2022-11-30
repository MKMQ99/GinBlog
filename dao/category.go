package dao

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"gorm.io/gorm"
)

// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate model.Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// CreateCate 新增分类
func CreateCate(data *model.Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetCateInfo 查询单个分类信息
func GetCateInfo(id int) (model.Category, int) {
	var cate model.Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCSE
}

// GetCate 查询分类列表
func GetCate(pageSize int, pageNum int) ([]model.Category, int64) {
	var cate []model.Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCate 编辑分类信息
func EditCate(id int, data *model.Category) int {
	var cate model.Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate model.Category
	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
