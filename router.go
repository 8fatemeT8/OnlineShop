package main

import (
	"OnlineShop/config"
	"OnlineShop/controller"
	_ "OnlineShop/docs"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

var router = gin.Default()
var public = router.Group("/auth")
var protected = router.Group("/api")
var adminRoute = router.Group("/api/admin")

func SetRouters() {
	SetSwagger()

	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)

	protected.GET("/product", controller.GetProducts)
	protected.GET("/product/menu", controller.GetMenu)

	adminRoute.GET("/say-hi", controller.AdminStaff)

	router.Run(":8090")
}

func SetSwagger() {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func ManageOriginsMiddleware() {
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8090", "http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}))

	public.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8090", "http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: false,
		Debug:            true,
	}))

	protected.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8090", "http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	}))

	adminRoute.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8090", "http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	}))
}

func SetMiddleware() {
	ManageOriginsMiddleware()
	public.Use(config.ValidateRequest)
	protected.Use(config.JwtAuthMiddleware())
	adminRoute.Use(config.HasAdminAccess)
}
