package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique"`
	Password  string
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	RoleID    uint
	Role      Role
	Baskets   []Basket
}

type Role struct {
	gorm.Model
	Name string
	User []User
}

type Authentication struct {
	Username  string `json:"username"  binding:"required"`
	Password  string `json:"password"  binding:"required"`
	Firstname string `json:"firstname"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"unique" json:"email"`
}

type LoginDto struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type Token struct {
	Role        string `json:"role"`
	Username    string `json:"username"`
	TokenString string `json:"token"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
