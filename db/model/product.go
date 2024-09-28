package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id             uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name           string
	Description    string
	Price          float64
	ProductDetails []ProductDetail
	Comments       *[]Comment
	Categories     []Category `gorm:"many2many:product_categories;"`
}

type ProductDetail struct {
	gorm.Model
	Id        uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	ColorId   uint
	Color     Color
	SizeId    uint
	Size      Size
	ProductId uint
	Product   Product
	Count     uint
}

type Color struct {
	gorm.Model
	Id   uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name string
}

type Size struct {
	gorm.Model
	Id   uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name string
}

type Category struct {
	gorm.Model
	Id   uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name string
}

type Comment struct {
	gorm.Model
	Id        uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	ProductId uint
	Product   Product
	UserId    uint
	User      User
	Message   string
}

type Favorite struct {
	gorm.Model
	Id        uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	ProductId uint
	UserId    uint
}
