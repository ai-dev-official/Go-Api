package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/aidev/crud-in-go/initializers"
	"golang.org/aidev/crud-in-go/models"
	"gorm.io/gorm"
)

// Define a struct to parse incoming request body for categories
var categoryBody struct {
	Name string
}




func CreateCategory(c *gin.Context) {

	c.Bind(&categoryBody)

	// Create new category instance
	category := models.Category{
		Name: categoryBody.Name,
	}

	// Insert category into the database
	result := initializers.DB.Create(&category)

	// Check for errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Category created successfully",
		"category": category,
	})
}


// GetCategories retrieves all Categories from the database and returns them in the response
func GetCategories(c *gin.Context) {
	var categories []models.Category

	// Retrieve categories from the database
	result := initializers.DB.Find(&categories)

	// Check for errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return success response with categories
	c.JSON(http.StatusOK, gin.H{
		"message":  "Categories retrieved successfully",
		"categories": categories,
	})
}


// GetCategoryById retrieves a single category by its ID from the database
func GetCategoryById(c *gin.Context) {
	// Extract the category ID from the request parameters
	id := c.Param("id")

	var category models.Category

	// Retrieve the category from the database
	result := initializers.DB.First(&category, id)

	// Check for errors
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Category not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
		}
		return
	}

	// Return success response with the category
	c.JSON(http.StatusOK, gin.H{
		"message": "Category retrieved successfully",
		"Category": category,
	})
}


// Updatecategory updates an existing category by its ID
func UpdateCategory(c *gin.Context) {
	// Extract the category ID from the request parameters
	id := c.Param("id")

	c.Bind(&categoryBody)

	// Print received data for debugging purposes
	// Remove these in categoryion code
	fmt.Printf("Received data: %+v\n", categoryBody)

	// Find the category by ID
	var category models.Category
	result := initializers.DB.First(&category, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Category not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
		}
		return
	}

	// Update the Category fields

	initializers.DB.Model(&category).Updates(models.Category{
		Name: categoryBody.Name,
	})

	// Save the updated Category
	result = initializers.DB.Save(&category)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Return success response with the updated Category
	c.JSON(http.StatusOK, gin.H{
		"message": "Category updated successfully",
		"category": category,
	})
}


// DeleteCategory deletes a Category by its ID
func DeleteCategory(c *gin.Context) {
	// Extract the Category ID from the request parameters
	id := c.Param("id")

	// Delete the Category from the database
	result := initializers.DB.Delete(&models.Category{}, id)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Category not found",
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
		"message": "Category deleted successfully",
	})
}

