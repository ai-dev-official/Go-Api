package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/aidev/crud-in-go/controllers"
	"golang.org/aidev/crud-in-go/initializers"
)

func init() {
	initializers.LoadVariables()
	initializers.DbConnect()

}

func main() {
	r := gin.Default()

	// Product routes
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	// Category routes
	r.POST("/categories", controllers.CreateCategory)
	r.GET("/categories", controllers.GetCategories)
	r.GET("/categories/:id", controllers.GetCategoryById)
	r.PUT("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/categories/:id", controllers.DeleteCategory)

	r.Run()
}







