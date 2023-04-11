package controllers

import (
	"go-jwt/database"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    Product,
	})
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"data":    Product,
	})
}

func FindProducts(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userRole := userData["role"]
	userID := uint(userData["id"].(float64))

	var products []models.Product

	// Jika userRole = "admin", maka cari semua produk tanpa memperdulikan userID
	if userRole == "admin" {
		err := db.Debug().Find(&products).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to retrieve products",
				"error":   err.Error(),
			})
			return
		}
	} else { // Jika userRole bukan "admin", maka cari produk berdasarkan userID
		err := db.Debug().Where("user_id = ?", userID).Find(&products).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to retrieve products",
				"error":   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Products retrieved successfully",
		"data":    products,
	})
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	userRole := userData["role"]
	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete product",
			"error":   "Invalid product ID",
		})
		return
	}

	var product models.Product
	// Jika userRole = "admin", maka cari product tanpa memperdulikan userID
	if userRole == "admin" {
		err = db.Debug().Where("id = ?", productId).First(&product).Error
	} else {
		err = db.Debug().Where("id = ? AND user_id = ?", productId, userID).First(&product).Error
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete product",
			"error":   "Product not found or does not belong to the user",
		})
		return
	}

	err = db.Debug().Delete(&product).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})

}

func GetProductByID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	userRole := userData["role"]
	productId, err := strconv.Atoi(c.Param("productId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to retrieve product",
			"error":   "Invalid product ID",
		})
		return
	}

	var product models.Product

	// Jika userRole = "admin", maka cari product tanpa memperdulikan userID
	if userRole == "admin" {
		err = db.Debug().Where("id = ?", productId).First(&product).Error
	} else {
		err = db.Debug().Where("id = ? AND user_id = ?", productId, userID).First(&product).Error
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to retrieve product",
			"error":   "Product not found or does not belong to the user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product retrieved successfully",
		"data":    product,
	})
}
