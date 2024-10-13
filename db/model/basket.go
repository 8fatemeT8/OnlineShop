package model

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	Address      string
	UserID       uint
	User         User
	Transactions []Transaction
	LineItems    []LineItem
}

type LineItem struct {
	gorm.Model
	ProductPrice    float64
	ProductDiscount float64
	Description     string
	ProductDetailID uint
	ProductDetail   ProductDetail
	BasketId        uint
}
