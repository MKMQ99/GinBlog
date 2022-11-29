package main

import (
	"GinBlog/model"
	"GinBlog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
