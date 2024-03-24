package main

import (
	"OnlineShop/config"
	"OnlineShop/controller"
	model "OnlineShop/db"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"net/http"
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

func ManageOriginsMiddleware() {
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}))
	public.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}))
	protected.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}))
	adminRoute.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}))
}

func SetMiddleware() {
	ManageOriginsMiddleware()
	public.Use(config.ValidateRequest)
	protected.Use(config.JwtAuthMiddleware())
	adminRoute.Use(config.HasAdminAccess)
}

func main() {
	model.CreateTables()
	SetMiddleware()
	SetRouters()
}
