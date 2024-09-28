package main

import (
	model "OnlineShop/db"
	_ "OnlineShop/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
func main() {
	model.CreateTables()
	SetMiddleware()
	SetRouters()
}
