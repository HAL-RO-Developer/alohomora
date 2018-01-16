package controller

import "github.com/gin-gonic/gin"

func Hoge(c *gin.Context) {
	c.JSON(200, gin.H{
		"hoge": "hoge",
	})
}

func hoge(c *gin.Context) {
	c.JSON(200, gin.H{
		"hoge": "hoge",
	})
}
