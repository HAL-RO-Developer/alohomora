package controller

import (
	"net/http"

	"github.com/HAL-RO-Developer/alohomora/model"
	"github.com/HAL-RO-Developer/alohomora/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := model.User{}
	user.Email = c.PostForm("email")
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	ok := service.User.ExisByEmail(user.Email)
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "メールが重複しています",
		})
		return
	}
	ok = service.User.ExisByName(user.Name)
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "名前が重複しています",
		})
		return
	}
	user = service.User.Store(user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user, ok := service.User.Login(email, password)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "ログインエラー",
		})
		return
	}
	token := service.Token.Create(user.ID)
	c.JSON(http.StatusOK, token)
}
