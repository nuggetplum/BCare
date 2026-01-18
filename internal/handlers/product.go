package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuggetplum/BCare/internal/database"
	"github.com/nuggetplum/BCare/internal/models"
)

// GET /products
func GetProducts(c *gin.Context) {
	var products []models.Product
	// Select * from products
	result := database.DB.Find(&products)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GET /products/:id
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	// Select * from products where id = ?
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
