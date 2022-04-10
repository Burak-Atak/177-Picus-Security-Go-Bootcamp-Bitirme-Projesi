package cart

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetCartList(userId uint) ([]CartDetails, error) {
	allCartDetails := GetAllCartDetailsOfUser(userId)
	if len(allCartDetails) == 0 {
		return nil, helpers.CartIsEmptyError
	}
	return allCartDetails, nil
}

func (s *Service) AddProductToCart(productId uint, amount int, userId uint) (*product.Product, error) {

	chosenProduct := product.SearchById(productId)
	if chosenProduct.ID == 0 {
		return nil, helpers.ProductNotFoundError
	}

	if chosenProduct.Stock <= amount {
		return nil, helpers.ProductNotEnoughStockError
	}

	if chosenProduct.Stock <= 0 {
		return nil, helpers.InvalidNumberOfProductsError
	}

	if IsProductExist(userId, productId) {
		return nil, helpers.ProductAlreadyExistInCart
	}

	return chosenProduct, nil
}

// UpdateProductInCart updates the amount of a product in the cart
func (s *Service) UpdateProductInCart(productId uint, userId uint, amount int) (*CartDetails, error) {
	chosenProduct := product.SearchById(productId)
	if chosenProduct.ID == 0 {
		return nil, helpers.ProductNotFoundError
	}

	if chosenProduct.Stock <= amount {
		return nil, helpers.ProductNotEnoughStockError
	}

	if chosenProduct.Stock <= 0 {
		return nil, helpers.InvalidNumberOfProductsError
	}
	if !IsProductExist(userId, productId) {
		return nil, helpers.ProductNotFoundErrorInCart
	}

	cartDetails := GetCartDetailsByCartIdAndProductId(userId, productId)
	return cartDetails, nil
}

func (s *Service) DeleteProductFromCart(userId, productId uint) (*CartDetails, error) {
	chosenProduct := product.SearchById(productId)
	if chosenProduct.ID == 0 {
		return nil, helpers.ProductNotFoundError
	}
	if !IsProductExist(userId, productId) {
		return nil, helpers.ProductNotFoundErrorInCart
	}
	cartDetails := GetCartDetailsByCartIdAndProductId(userId, productId)
	return cartDetails, nil
}

func (s *Service) GetCartsHasProduct(productId uint) (*[]CartDetails, error) {
	allCartDetails := GetCartDetailsByProductId(productId)
	if len(*allCartDetails) == 0 {
		return nil, helpers.CartIsEmptyError
	}
	return allCartDetails, nil
}
