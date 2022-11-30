package v1

import (
	"GinBlog/dao"
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 查询用户是否存在
func UerExist(c *gin.Context) {
	//
}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBind(&data)
	code = dao.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		dao.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data := dao.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&data)

	code := dao.CheckUpUser(id, data.Username)
}

// 删除用户
func DeleteUser(c *gin.Context) {

}
