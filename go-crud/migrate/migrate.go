package main

import (
	"crud/initializers"
	"crud/models"
)

func init() {
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Products{})
}
