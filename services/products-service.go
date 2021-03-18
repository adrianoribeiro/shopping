package services

import (
	"github.com/shopping/repo/database"
	"github.com/shopping/repo/models"
)

func NewProduct(product *models.Product) error {

	database.DB.Create(&product)

	//TODO Check it
	return nil
}
func GetProduct(id int) (bool, models.Product) {

	var product models.Product
	database.DB.Find(&product, id)

	if product.Id == 0 {
		return false, models.Product{}
	}

	return true, product
}

func GetProducts() []models.Product {
	var products []models.Product
	database.DB.Find(&products)

	return products
}

//TODO I'll check this implementation it not sounds good
func Delete(id int) bool{
	found, product := GetProduct(id)

	if !found {
		return false
	}
	database.DB.Delete(product)
	return true
}