package main

import "github.com/makki0205/template/model"

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.Receipt{})
	db.DropTableIfExists(&model.Sales{})
	db.DropTableIfExists(&model.Seller{})
	db.DropTableIfExists(&model.Product{})
	db.DropTableIfExists(&model.ProductImages{})

	db.AutoMigrate(&model.Receipt{})
	db.AutoMigrate(&model.Sales{})
	db.AutoMigrate(&model.Seller{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.ProductImages{})
}
