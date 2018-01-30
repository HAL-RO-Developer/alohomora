package service

import (
	"fmt"
	"testing"
)

func TestUser_Login(t *testing.T) {

	ok, _ := User.Login("llxo2_5oxll@icloud.com3", "hoge")
	fmt.Println(ok)
}