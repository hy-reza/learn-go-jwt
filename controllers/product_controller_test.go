package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProductByIDFound(t *testing.T) {
	productServiceMock := &ProductServiceMock{}
	productController := ProductController{productService: productServiceMock}

	router := gin.Default()

	router.GET("/products/:productId", func(c *gin.Context) {
		c.Set("userData", jwt.MapClaims{
			"id":   float64(1),
			"role": "user",
		})
		productController.GetProductByID(c)
	})

	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetProductByIDNotFound(t *testing.T) {
	productServiceMock := &ProductServiceMock{}
	productController := ProductController{productService: productServiceMock}

	router := gin.Default()

	router.GET("/products/:productId", func(c *gin.Context) {
		c.Set("userData", jwt.MapClaims{
			"id":   float64(1),
			"role": "user",
		})
		productController.GetProductByID(c)
	})

	req, _ := http.NewRequest(http.MethodGet, "/products/999", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestFindProductsAsAdmin(t *testing.T) {
	productServiceMock := &ProductServiceMock{}
	productController := ProductController{productService: productServiceMock}

	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		c.Set("userData", jwt.MapClaims{
			"id":   float64(1),
			"role": "admin",
		})
		productController.FindProducts(c)
	})

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestFindProductsAsUser(t *testing.T) {
	productServiceMock := &ProductServiceMock{}
	productController := ProductController{productService: productServiceMock}

	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		c.Set("userData", jwt.MapClaims{
			"id":   float64(1),
			"role": "user",
		})
		productController.FindProducts(c)
	})

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
