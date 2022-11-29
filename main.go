package main

import (
	"GinBlog/dao"
	"GinBlog/routes"
)

func main() {
	dao.InitDb()
	routes.InitRouter()
}
