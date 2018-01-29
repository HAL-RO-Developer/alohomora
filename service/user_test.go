package service

import (
	"fmt"
	"testing"
)

func TestUser_Login(t *testing.T) {

	ok := User.Login("llxo2_5oxll@icloud.com3", "hoge")
	fmt.Println(ok)
}
