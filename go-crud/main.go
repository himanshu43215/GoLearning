package main

import (
	"crud/controllers"
	"crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.POST("v1/addProduct", controllers.ProductCreate)
	r.GET("v1/Products", controllers.GetAllProducts)
	r.GET("/Products/:id", controllers.GetProductById)
	r.PUT("/Products/:id", controllers.UpdateProductDetails)
	r.DELETE("/Products/:id", controllers.DeleteProduct)

	r.Run()
}
