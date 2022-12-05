package dao

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"gorm.io/gorm"
)

// AddComment 新增评论
func AddComment(data *model.Comment) int {
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetComment 查询单个评论
func GetComment(id int) (model.Comment, int) {
	var comment model.Comment
	err := db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return comment, errmsg.ERROR
	}
	return comment, errmsg.SUCCSE
}

// GetCommentList 后台所有获取评论列表
func GetCommentList(pageSize int, pageNum int) ([]model.Comment, int64, int) {
	var comments []model.Comment
	var total int64
	db.Find(&comments).Count(&total)
	err = db.Model(&comments).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&comments).Error
	if err != nil {
		return comments, 0, errmsg.ERROR
	}
	return comments, total, errmsg.SUCCSE
}

// GetCommentCount 获取评论数量
func GetCommentCount(id int) int64 {
	var comment model.Comment
	var total int64
	db.Find(&comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	return total
}

// GetCommentListFront 展示页面获取评论列表
func GetCommentListFront(id int, pageSize int, pageNum int) ([]model.Comment, int64, int) {
	var commentList []model.Comment
	var total int64
	db.Find(&model.Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	err = db.Model(&model.Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?",
		id).Where("status = ?", 1).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCSE
}

// DeleteComment 删除评论
func DeleteComment(id uint) int {
	var comment model.Comment
	err = db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// CheckComment 通过评论
func CheckComment(id int, data *model.Comment) int {
	var comment model.Comment
	var res model.Comment
	var article model.Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// UncheckComment 撤下评论
func UncheckComment(id int, data *model.Comment) int {
	var comment model.Comment
	var res model.Comment
	var article model.Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
