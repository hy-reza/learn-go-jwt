package controllers

import (
	"errors"
	"go-jwt/models"
)

type ProductServiceMock struct{}

func (psm *ProductServiceMock) GetProductByID(userID uint, userRole string, productID int) (models.Product, error) {
	if productID == 1 {
		return models.Product{
			ID:          1,
			Title:       "Sample Product",
			Description: "Sample Product Description",
			UserID:      userID,
		}, nil
	} else {
		return models.Product{}, errors.New("Product not found")
	}
}

func (psm *ProductServiceMock) FindProducts(userID uint, userRole string) ([]models.Product, error) {
	return []models.Product{
		{
			ID:          1,
			Title:       "Sample Product 1",
			Description: "Sample Product Description 1",
			UserID:      userID,
		},
		{
			ID:          2,
			Title:       "Sample Product 2",
			Description: "Sample Product Description 2",
			UserID:      userID,
		},
	}, nil
}
