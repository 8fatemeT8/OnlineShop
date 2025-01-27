package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name           string
	Description    string
	Price          float64
	ProductDetails []ProductDetail
	Comments       *[]Comment
	Categories     []Category `gorm:"many2many:product_categories;"`
}

type ProductDetail struct {
	gorm.Model
	ColorID   uint
	Color     Color
	SizeID    uint
	Size      Size
	ProductID uint
	Product   Product
	Count     uint
}

type Color struct {
	gorm.Model
	Name string
}

type Size struct {
	gorm.Model
	Name string
}

type Category struct {
	gorm.Model
	Name string
}

type Comment struct {
	gorm.Model
	ProductID uint
	Product   Product
	UserID    uint
	User      User
	Message   string
}

type Favorite struct {
	gorm.Model
	ProductId uint
	UserId    uint
}
