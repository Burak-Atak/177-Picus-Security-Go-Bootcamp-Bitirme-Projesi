package helpers

import "errors"

var (
	CartIsEmptyError             = errors.New("your cart is empty")
	InvalidIdError               = errors.New("invalid id")
	UserNotFoundError            = errors.New("user not found")
	UserExistsError              = errors.New("user already exists")
	InvalidPasswordError         = errors.New("invalid password")
	UnAuthorizedError            = errors.New("you are not authorized to access this resource")
	ProductNotFoundError         = errors.New("product not found")
	ProductNotFoundErrorInCart   = errors.New("product not found in cart")
	ProductNotEnoughStockError   = errors.New("product not enough stock")
	UnsupportedMediaType         = errors.New("Content-Type header is not application/json")
	InvalidNumberOfProductsError = errors.New("invalid number of products")
	ProductAlreadyExistError     = errors.New("product already exist")
	ProductAlreadyExistInCart    = errors.New("product already exist in cart")
	CategoryNotFoundError        = errors.New("category not found")
	CategoryAlreadyExistError    = errors.New("category already exist")
)
