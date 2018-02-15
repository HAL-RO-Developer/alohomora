package router

import (
	"github.com/HAL-RO-Developer/alohomora/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	//api.GET("/hoge", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"hoge": "hoge",
	//	})
	//})

	api.POST("/user", controller.CreateUser)
	api.POST("/login", controller.Login)
	api.POST("/uuid", controller.SetToken)
}
