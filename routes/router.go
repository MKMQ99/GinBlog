package routes

import (
	v1 "GinBlog/api/v1"
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		router.GET("admin/category", v1.GetCate)
		router.POST("category/add", v1.AddCategory)
		router.PUT("category/:id", v1.EditCate)
		router.DELETE("category/:id", v1.DeleteCate)
		// 文章模块的路由接口
		router.GET("admin/article/info/:id", v1.GetArtInfo)
		router.GET("admin/article", v1.GetArt)
		router.GET("admin/article/list/:id", v1.GetCateArt)
		router.POST("article/add", v1.AddArticle)
		router.PUT("article/:id", v1.EditArt)
		router.DELETE("article/:id", v1.DeleteArt)
	}
	r.Run(utils.HttpPort)
}
