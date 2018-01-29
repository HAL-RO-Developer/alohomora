package main

import "github.com/HAL-RO-Developer/alohomora/model"

func main() {
	db := model.GetDBConn()
	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.Key{})
	db.DropTableIfExists(&model.Token{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Key{})
	db.AutoMigrate(&model.Token{})
}
