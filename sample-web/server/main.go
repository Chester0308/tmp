package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sample-web/server/controller"
	"github.com/gin-contrib/multitemplate"
)

func main() {
	router := gin.Default()
	router.HTMLRender = createRender()

	router.GET("/", controller.GetLogin)
	router.POST("/login", controller.PostLogin)

	router.Run()
}

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "./templates/layouts/layout.html", "./templates/contents/index.html")
	r.AddFromFiles("login", "./templates/layouts/layout.html", "./templates/contents/login.html")
	return r
}