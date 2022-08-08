package services

import (
	"crud-go/database"
	"crud-go/models"
)

func CheckIfProductExists(Id string) bool {
	var product models.Product
	database.Instance.First(&product, Id)
	if product.ID == 0 {
		return false
	}
	return true
}
