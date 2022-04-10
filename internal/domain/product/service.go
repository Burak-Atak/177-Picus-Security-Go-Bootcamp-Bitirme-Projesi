package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// GetProduct checks if product exists and returns it
func (s *Service) GetProduct(productId uint) (*Product, error) {

	product := SearchById(productId)
	if product.ID == 0 {
		return nil, helpers.ProductNotFoundError
	}
	return product, nil

}

// CreateProduct checks if the product is already in the database
func (s *Service) CreateProduct(productName string, sku string) error {

	if IsProductExist(productName, sku) {
		return helpers.ProductAlreadyExistError
	}

	return nil

}

// GetProductList returns a list of products if there are any in the database otherwise returns error
func (s *Service) GetProductList() ([]Product, error) {
	products := FindAll()

	if len(products) == 0 {
		return nil, helpers.ProductNotFoundError
	}

	return products, nil
}

// SearchProduct checks if product is exist in database by product name and sku
func (s *Service) SearchProduct(search string) ([]Product, error) {
	products := SearchProduct(search)

	if len(products) == 0 {
		return nil, helpers.ProductNotFoundError
	}

	return products, nil
}
