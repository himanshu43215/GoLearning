package main

import (
	"bytes"
	"crud/controllers"
	"crud/initializers"
	"crud/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllProducts(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define the route using the same path and controller function
	router.GET("v1/Products", controllers.GetAllProducts)

	// Create a new HTTP request for the defined route
	req, err := http.NewRequest("GET", "/v1/Products", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request and record the response
	router.ServeHTTP(recorder, req)

	// Check the response status code
	//assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response struct {
		Products []models.Products `json:"products"`
	}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Add your assertions to validate the response
	// For example:
	//assert.Len(t, response.Products, 0) // Assert that there are no products in the response
}
func TestProductCreate(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Set up the database connection or any required initializers
	// Here, assuming you have an initializers package with DB setup
	//	initializers.SetupDB()
	initializers.ConnectToDB()
	// Define the route using the same path and controller function
	router.POST("v1/addProduct", controllers.ProductCreate)

	// Create a request body
	requestBody := []byte(`{
		"ProductName": "laptop",
		"TypeProduct": "Example Type",
		"ProductPrice": "10.99",
		"Quantity": "100"
	}`)

	// Create a new HTTP request for the defined route
	req, err := http.NewRequest("POST", "/v1/addProduct", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request and record the response
	router.ServeHTTP(recorder, req)

	// Check the response status code
	//assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response struct {
		Product models.Products `json:"product"`
	}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Add your assertions to validate the response
	// For example:
	//assert.Equal(t, "Example Product", response.Product.ProductName) // Assert that the product name matches the expected value
}
func TestGetProductById(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Set up the database connection or any required initializers
	// Here, assuming you have an initializers package with DB setup
	//initializers.SetupDB()
	initializers.ConnectToDB()
	// Define the route using the same path and controller function
	router.GET("/Products/:id", controllers.GetProductById)

	// Create a new HTTP request for the defined route
	req, err := http.NewRequest("GET", "/Products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request and record the response
	router.ServeHTTP(recorder, req)

	// Check the response status code
	//assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response struct {
		Product models.Products `json:"product"`
	}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Add your assertions to validate the response
	// For example:
	//assert.Equal(t, 1, response.Product.ID) // Assert that the product ID matches the expected value
}
func TestUpdateProductDetails(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Set up the database connection or any required initializers
	// Here, assuming you have an initializers package with DB setup
	initializers.ConnectToDB()

	// Define the route using the same path and controller function
	router.PUT("/Products/:id", controllers.UpdateProductDetails)

	// Create a request body
	requestBody := []byte(`{
		"ProductName": "Updated Product",
		"TypeProduct": "Updated Type",
		"ProductPrice": "19.99",
		"Quantity": "50"
	}`)

	// Create a new HTTP request for the defined route
	req, err := http.NewRequest("PUT", "/Products/1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request and record the response
	router.ServeHTTP(recorder, req)

	// Check the response status code
	//assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response struct {
		Product models.Products `json:"product"`
	}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Add your assertions to validate the response
	// For example:
	//assert.Equal(t, "Updated Product", response.Product.ProductName) // Assert that the product name is updated
	//assert.Equal(t, "Updated Type", response.Product.TypeProduct)    // Assert that the type is updated
	//assert.Equal(t, "19.99", response.Product.ProductPrice)          // Assert that the product price is updated
	//assert.Equal(t, "50", response.Product.Quantity)                 // Assert that the quantity is updated

}
func TestDeleteProduct(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Set up the database connection or any required initializers
	// Here, assuming you have an initializers package with DB setup
	initializers.ConnectToDB()

	// Define the route using the same path and controller function
	router.DELETE("/Products/:id", controllers.DeleteProduct)

	// Create a new HTTP request for the defined route
	req, err := http.NewRequest("DELETE", "/Products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request and record the response
	router.ServeHTTP(recorder, req)

	// Check the response status code
	//assert.Equal(t, http.StatusOK, recorder.Code)

	// Verify that the product is deleted from the database
	//var product models.Products
	//result := initializers.DB.First(&product, 1)
	//assert.Error(t, result.Error) // Assert that an error is returned, indicating that the product is not found
}
