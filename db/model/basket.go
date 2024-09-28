package model

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	Id           uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Address      string
	UserId       uint
	User         User
	Transactions []Transaction
	LineItems    []LineItem
}

type LineItem struct {
	gorm.Model
	Id              uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	ProductPrice    float64
	ProductDiscount float64
	Description     string
	ProductDetailId uint
	ProductDetail   ProductDetail
	BasketId        uint
}
