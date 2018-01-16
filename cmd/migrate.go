package main

import "github.com/HAL-RO-Developer/alohomora/model"

func main() {
	db := model.GetDBConn()
	db.AutoMigrate(&model.User{})
}
