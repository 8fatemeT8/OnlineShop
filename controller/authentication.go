package controller

import (
	"OnlineShop/config"
	model2 "OnlineShop/db"
	"OnlineShop/db/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var input model.Authentication
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var role model.Role
	model2.DB.Model(model.Role{}).Where(model.Role{Name: "user"}).Take(&role)
	user := initialUser(input)
	user.Role = role

	result := model2.DB.Model(&model.User{}).Create(&user)
	err := result.Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {

	var input model.LoginDto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := config.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func AdminStaff(c *gin.Context) {
	c.JSON(200, "hiii adminssssss")
}
func initialUser(input model.Authentication) model.User {
	u := model.User{}

	u.Username = input.Username
	u.Password = input.Password
	u.FirstName = input.Firstname
	u.LastName = input.LastName
	u.Email = input.Email
	return u
}
