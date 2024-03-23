package main

import (
	model2 "OnlineShop/db"
	"OnlineShop/db/model"
	"gorm.io/gen"
)

type Querier struct {
}

func main() {
	g := gen.NewGenerator(
		gen.Config{
			OutPath: "../OnlineShop/db/query",
			Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		})

	g.UseDB(model2.DB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {}, model.User{})

	// Generate the code
	g.Execute()
}
