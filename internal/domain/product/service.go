package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetProduct(productId uint) (*Product, error) {

	product := SearchById(productId)
	if product.ID == 0 {
		return nil, helpers.ProductNotFoundError
	}
	return product, nil

}
func (s *Service) CreateProduct(productName string, sku string) error {

	if IsProductExist(productName, sku) {
		return helpers.ProductAlreadyExistError
	}

	return nil

}

func (s *Service) GetProductList() ([]Product, error) {
	products := FindAll()

	if len(products) == 0 {
		return nil, helpers.ProductNotFoundError
	}

	return products, nil
}

func (s *Service) SearchProduct(search string) ([]Product, error) {
	products := SearchProduct(search)

	if len(products) == 0 {
		return nil, helpers.ProductNotFoundError
	}

	return products, nil
}
