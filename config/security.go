package config

import (
	model2 "OnlineShop/db"
	"OnlineShop/db/model"
	"golang.org/x/crypto/bcrypt"
)

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := model.User{}

	err = model2.DB.Model(model.User{}).Preload("Role").Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = CheckPasswordHash(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := GenerateJWT(u.Username, u.Role.Name)

	if err != nil {
		return "", err
	}

	return token, nil

}
