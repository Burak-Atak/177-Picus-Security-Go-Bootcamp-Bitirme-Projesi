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
func (s *Service) GetProductList(pageIndex int, pageSize int) ([]Product, int) {
	products, allProducts := GetAll(pageIndex, pageSize)
	return products, allProducts
}

// SearchProduct checks if product is exist in database by product name and sku
func (s *Service) SearchProduct(searchQuery string) ([]Product, error) {
	products := SearchProduct(searchQuery)

	if len(products) == 0 {
		return nil, helpers.ProductNotFoundError
	}

	return products, nil
}

func (s *Service) SearchProductWithPagination(searchQuery string, pageIndex int, pageSize int) []Product {
	products := SearchProductWithPagination(searchQuery, pageIndex, pageSize)
	return products
}
