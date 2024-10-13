package controller

import (
	"OnlineShop/config"
	model2 "OnlineShop/db"
	"OnlineShop/db/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register
// PingExample godoc
// @Summary ping example
// @Schemes
// @Param  input body  model.Authentication true "User Data"
// @Description do ping
// @Tags UserManagement
// @Accept json
// @Produce json
// @Success 200
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var input model.Authentication
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var role model.Role
	model2.Db.Model(model.Role{}).Where(model.Role{Name: "user"}).Take(&role)
	user := initialUser(input)
	user.Role = role

	result := model2.Db.Model(&model.User{}).Create(&user)
	err := result.Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := config.LoginCheck(input.Username, input.Password)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Login
// PingExample godoc
// @Summary ping example
// @Schemes
// @Param  input body  model.LoginDto true "Login data"
// @Description do ping
// @Tags UserManagement
// @Accept json
// @Produce json
// @Success 200
// @Router /auth/login [post]
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

// GetV2UserByUsername
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags UserManagement
// @Accept json
// @Produce json
// @Success 200
// @Router /api/admin/v2/user/{username} [get]
func GetV2UserByUsername(c *gin.Context) {
	var username = c.Param("username")
	var user = model.User{}
	model2.Db.Model(&model.User{}).Where(&model.User{Username: username}).First(&user)
	c.JSON(http.StatusOK, user)
}

// AdminStaff
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags Admin
// @Accept json
// @Security
// @Produce json
// @Success 200
// @Router /api/admin/say-hi [get]
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
