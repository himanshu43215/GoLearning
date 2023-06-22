package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ProductName  string `json:"Product Name"`
	TypeProduct  string `json:"Product Description"`
	ProductPrice string `json:"Product Price"`
	Quantity     string `json:"Quantity"`
}
