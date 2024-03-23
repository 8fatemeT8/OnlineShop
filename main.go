package main

import (
	"OnlineShop/config"
	"OnlineShop/controller"
	model "OnlineShop/db"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()
var public = router.Group("/auth")
var protected = router.Group("/api")
var adminRoute = router.Group("/api/admin")

func SetRouters() {
	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)
	protected.GET("/hi", func(context *gin.Context) {
		context.JSON(200, "hiiiii")
	})

	adminRoute.GET("/say-hi", controller.AdminStaff)
	router.Run("localhost:8085")
}

func SetMiddleware() {
	// Middleware 1 - Request validation
	public.Use(config.ValidateRequest)
	protected.Use(config.JwtAuthMiddleware())
	adminRoute.Use(config.HasAdminAccess)
}

func main() {
	model.CreateTables()
	SetMiddleware()
	SetRouters()
}
