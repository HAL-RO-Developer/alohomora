package router

import (
	"github.com/HAL-RO-Developer/alohomora/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/uuid", controller.SetToken)
	api.POST("/open", controller.Open)
	api.POST("/close", controller.Close)
	api.POST("/device", controller.SetDevice)
}
