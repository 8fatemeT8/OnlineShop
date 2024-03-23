package model

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	Address      string
	UserId       uint
	User         User
	Transactions []Transaction
	LineItems    []LineItem
}

type LineItem struct {
	gorm.Model
	ProductPrice    float64
	ProductDiscount float64
	Description     string
	ProductDetailId uint
	ProductDetail   ProductDetail
	BasketId        uint
}
