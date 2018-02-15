package router

import (
	"github.com/HAL-RO-Developer/alohomora/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/user", controller.CreateUser)
	api.POST("/login", controller.Login)
	api.POST("/uuid", controller.SetToken)
}
