package routes

import (
	"fmt"
	"loghawk/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductRoutes(db *gorm.DB, router *gin.Engine) {

	router.POST("/products", func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, product)
	})

	// Read All (GET)
	router.GET("/products", func(c *gin.Context) {
		var products []models.Product
		db.Find(&products)
		c.JSON(http.StatusOK, products)
	})

	// Read One (GET)
	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		db.First(&product, "id = ?", id)

		if product.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}

		c.JSON(http.StatusOK, product)
	})

	// Update (PUT)
	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		_id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Failed to delete logs.")
		}
		var updatedProduct models.Product
		var existingProduct models.Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("Updated Product > ", updatedProduct)
		db.First(&existingProduct, _id)
		fmt.Println("existing Product > ", existingProduct)
		db.Model(&models.Product{ID: uint(_id)}).Updates(&updatedProduct)

		c.JSON(http.StatusOK, updatedProduct)
	})

	// Delete (DELETE)
	router.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		_id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Failed to delete logs.")
		}
		db.Delete(&models.Product{ID: uint(_id)})

		c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
	})
}
