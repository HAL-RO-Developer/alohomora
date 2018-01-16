package model

import "testing"

func TestStore(t *testing.T) {
	user := User{
		Name:     "hoge",
		Email:    "llxo2_5oxll@icloud.com",
		Password: "hoge",
	}
	UserStore(user)
}
