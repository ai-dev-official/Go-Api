package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/aidev/crud-in-go/initializers"
	"golang.org/aidev/crud-in-go/models"
	"gorm.io/gorm"
)

// Define a struct to parse incoming request body
var body struct {
	Name          string
	Brand         string
	CategoryID    uint
	Price         float64
	StockQuantity int
	Description   string
	ReleaseDate   time.Time
}


func CreateProduct(c *gin.Context) {

	c.Bind(&body)

	// Create new Product instance
	product := models.Product{
		Name:          body.Name,
		Brand:         body.Brand,
		CategoryID:    body.CategoryID,
		Price:         body.Price,
		StockQuantity: body.StockQuantity,
		Description:   body.Description,
		ReleaseDate:   body.ReleaseDate,
	}

	// Insert product into the database
	result := initializers.DB.Create(&product)

	// Check for errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}


// GetProducts retrieves all products from the database and returns them in the response
func GetProducts(c *gin.Context) {
	var products []models.Product

	// Retrieve products from the database
	result := initializers.DB.Find(&products)

	// Check for errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return success response with products
	c.JSON(http.StatusOK, gin.H{
		"message":  "Products retrieved successfully",
		"products": products,
	})
}


// GetProductById retrieves a single product by its ID from the database
func GetProductById(c *gin.Context) {
	// Extract the product ID from the request parameters
	id := c.Param("id")

	var product models.Product

	// Retrieve the product from the database
	result := initializers.DB.First(&product, id)

	// Check for errors
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
		}
		return
	}

	// Return success response with the product
	c.JSON(http.StatusOK, gin.H{
		"message": "Product retrieved successfully",
		"product": product,
	})
}


// UpdateProduct updates an existing product by its ID
func UpdateProduct(c *gin.Context) {
	// Extract the product ID from the request parameters
	id := c.Param("id")

	c.Bind(&body)

	// Print received data for debugging purposes
	// Remove these in production code
	fmt.Printf("Received data: %+v\n", body)

	// Find the product by ID
	var product models.Product
	result := initializers.DB.First(&product, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
		}
		return
	}

	// Update the product fields

	initializers.DB.Model(&product).Updates(models.Product{
		Name:          body.Name,
		Brand:         body.Brand,
		CategoryID:    body.CategoryID,
		Price:         body.Price,
		StockQuantity: body.StockQuantity,
		Description:   body.Description,
		ReleaseDate:   body.ReleaseDate,
	})

	// Save the updated product
	result = initializers.DB.Save(&product)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return success response with the updated product
	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": product,
	})
}


// DeleteProduct deletes a product by its ID
func DeleteProduct(c *gin.Context) {
	// Extract the product ID from the request parameters
	id := c.Param("id")

	// Delete the product from the database
	result := initializers.DB.Delete(&models.Product{}, id)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
		}
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
