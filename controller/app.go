package controller

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var uid = ""

func SetToken(c *gin.Context) {
	token := c.Query("token")
	if token != "BAC92613-0C14-4F78-9C1E-93D431AF6870" {
		c.JSON(400, "token err")
		return
	}
	uid = uuid.NewV4().String()
	c.JSON(200, gin.H{
		"uuid": uid,
	})
}

func App(c *gin.Context) {
	uuid := c.Param("uid")
	if uid != uuid {
		c.Redirect(301, "https://konojunya.com")
		return
	}
	c.HTML(200, "index.html", nil)
}
