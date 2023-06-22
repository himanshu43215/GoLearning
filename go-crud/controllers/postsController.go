package controllers

import (
	"crud/initializers"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func ProductCreate(c *gin.Context) {
	//get the data off req boby
	var body struct {
		ProductName  string
		TypeProduct  string
		ProductPrice string
		Quantity     string
	}
	c.Bind(&body)

	//create the post
	product := models.Products{ProductName: body.ProductName, TypeProduct: body.TypeProduct, ProductPrice: body.ProductPrice, Quantity: body.Quantity}
	result := initializers.DB.Create(&product)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"product": product,
	})
}
func GetAllProducts(c *gin.Context) {
	//get the post
	var products []models.Products
	initializers.DB.Find(&products)
	//response with them
	c.JSON(200, gin.H{
		"products": products,
	})
}
func GetProductById(c *gin.Context) {
	//get id  off url
	id := c.Param("id")
	//get the post
	var product models.Products
	initializers.DB.First(&product, id)
	//response with them
	c.JSON(200, gin.H{
		"product": product,
	})
}
func UpdateProductDetails(c *gin.Context) {
	//Get the id of url
	id := c.Param("id")
	//get data of request bady
	var body struct {
		ProductName  string
		TypeProduct  string
		ProductPrice string
		Quantity     string
	}
	c.Bind(&body)
	//find the post
	var product models.Products
	initializers.DB.First(&product, id)
	//updated it
	initializers.DB.Model(&product).Updates(models.Products{
		ProductName:  body.ProductName,
		TypeProduct:  body.TypeProduct,
		ProductPrice: body.ProductPrice,
		Quantity:     body.Quantity,
	})
	//response it
	c.JSON(200, gin.H{
		"post": product,
	})

}
func DeleteProduct(c *gin.Context) {
	//get the od of the url
	id := c.Param("id")
	//delete the post
	initializers.DB.Delete(&models.Products{}, id)
	//response
	c.Status(200)

}
