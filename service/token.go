package service

import (
	"github.com/HAL-RO-Developer/alohomora/model"
	uuid "github.com/satori/go.uuid"
)

var Token = tokenimple{}

type tokenimple struct {
}

func (t *tokenimple) Create(userID uint) model.Token {
	uid := uuid.NewV4().String()
	to := model.Token{
		UserID: userID,
		Body:   uid,
	}
	db.Create(&to)
	return to
}
func (t *tokenimple) FindByToken(tokens string) (model.Token, bool) {
	var token model.Token
	db.Where("body = ?", tokens).First(&token)
	return token, token.UserID != 0
}
