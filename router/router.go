package router

import (
	"github.com/HAL-RO-Developer/alohomora/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")

	r.GET("/@:uid", controller.App)

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin.html", nil)
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	api := r.Group("/api")
	apiRouter(api)
	return r

}
