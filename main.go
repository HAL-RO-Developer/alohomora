package main

import "github.com/HAL-RO-Developer/alohomora/router"

func main() {
	r := router.GetRouter()
	r.Run(":3000")
}
