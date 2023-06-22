package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type ProductInventory struct {
	ProductName  string `json:"Product Name"`
	TypeProduct  string `json:"Product Description"`
	ProductPrice string `json:"Product Price"`
	Quantity     string `json:"Quantity"`
}
type MyResponse struct {
	Massage string `json:"Answer:"`
}

func HandleLambdaEvent(event ProductInventory) (MyResponse, error) {
	return MyResponse{Massage: fmt.Sprintf("%s is %d ", event.ProductName, event.TypeProduct, event.ProductPrice, event.Quantity)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
