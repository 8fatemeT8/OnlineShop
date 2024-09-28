package controller

import "github.com/gin-gonic/gin"

// GetProducts
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Security  JWT
// @Produce json
// @Success 200
// @Router /api/product [get]
func GetProducts(c *gin.Context) {
	response := []interface{}{
		map[string]any{
			"id":     1,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  "1,000,000",
		},
		map[string]any{
			"id":     12,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  "2,000,000",
		},
		map[string]any{
			"id":     13,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  "3,000,000",
		},
		map[string]any{
			"id":     14,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  "4,000,000",
		},
		map[string]any{
			"id":     15,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  "5,000,000",
		},
		map[string]any{
			"id":     16,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  "6,000,000",
		},
		map[string]any{
			"id":     17,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  "7,000,000",
		},
		map[string]any{
			"id":     18,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     19,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     2,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     21,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     22,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     23,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     24,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     25,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     26,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     27,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     3,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     28,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
		map[string]any{
			"id":     29,
			"title":  "title: Example",
			"name":   "Product name",
			"detail": "product detail product detail product detail product detail product detail",
			"price":  4000,
		},
	}
	c.JSON(200, response)
}

// GetMenu
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Security JWT
// @Produce json
// @Success 200
// @Router /api/product/menu [get]
func GetMenu(c *gin.Context) {
	response := []interface{}{
		map[string]any{
			"title": "Brand",
			"type":  "SELECT",
			"list":  []interface{}{"Prada", "ZARA", "Victoria Secret", "Prada2", "ZARA2", "Victoria Secret2"},
		},
		map[string]any{
			"title": "Category",
			"type":  "SELECT",
			"list":  []interface{}{"Prada", "ZARA", "Victoria Secret", "Prada", "ZARA3", "Victoria Secret", "Prada1", "ZARA2", "Victoria Secret1", "Prada"},
		},
		map[string]any{
			"title": "Color",
			"type":  "SELECT",
			"list":  []interface{}{"Prada", "ZARA", "Victoria Secret", "Prada", "ZARA1", "Victoria Secret1", "ZARA2", "Victoria Secret2", "ZARA3", "Victoria Secret3"},
		},
		map[string]any{
			"title": "Price",
			"type":  "SLIDER",
			"max":   "10000000",
			"min":   "300000",
		},
		map[string]any{
			"title": "Fabric material",
			"type":  "SELECT",
			"list":  []interface{}{"Prada", "ZARA4", "Victoria Secret", "Prada"},
		},
		map[string]any{
			"title": "Stock",
			"type":  "SELECT",
			"list":  []interface{}{"Prada", "ZARA2", "Victoria Secret", "Prada"},
		},
	}

	c.JSON(200, response)
}
