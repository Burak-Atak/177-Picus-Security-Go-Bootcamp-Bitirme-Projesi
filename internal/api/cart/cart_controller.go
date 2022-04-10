package cart

import (
	"fmt"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/cart"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	CartService *cart.Service
}

func NewCartController(cartService *cart.Service) *Controller {
	return &Controller{
		CartService: cartService,
	}
}

// GetCartList returns user's Product Cart List
func (c *Controller) GetCartList(context *gin.Context) {
	decodedToken, err := jwt.VerifyToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	allCartDetails, err := c.CartService.GetCartList(decodedToken.UserId)
	if err != nil {
		context.JSON(http.StatusAccepted, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	outPut := make([]Product, len(*allCartDetails))
	var TotalPrice float64
	for i, cartDetail := range *allCartDetails {
		TotalPrice += cartDetail.TotalPrice
		outPut[i] = Product{
			ProductName: cartDetail.ProductName,
			Amount:      cartDetail.Amount,
			UnitPrice:   cartDetail.UnitPrice,
			TotalPrice:  cartDetail.TotalPrice,
			ProductId:   cartDetail.ProductId,
		}
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success", "Total Price": TotalPrice, "data": outPut})

}

// DeleteProductFromCart deletes product from cart
func (c *Controller) DeleteProductFromCart(context *gin.Context) {
	decodedToken, err := jwt.VerifyToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	id, isOk := context.GetQuery("id")
	if !isOk {
		context.JSON(http.StatusBadRequest, gin.H{"message": helpers.IdIsRequiredError.Error()})
		context.Abort()
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": helpers.InvalidIdError.Error()})
		context.Abort()
		return
	}

	cartDetails, err := c.CartService.DeleteProductFromCart(decodedToken.UserId, uint(productId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	newAmount := -cartDetails.Amount
	newTotalPrice := -cartDetails.TotalPrice
	// Update user main cart
	cart.UpdateUserCart(decodedToken.UserId, newAmount, newTotalPrice)

	// Delete product from cart
	cart.DeleteProductInCart(decodedToken.UserId, uint(productId))

	context.JSON(http.StatusOK, gin.H{"message": "Product deleted from cart"})
}

// UpdateProductInCart updates product in cart
func (c *Controller) UpdateProductInCart(context *gin.Context) {
	var body RequestBody
	decodedToken, err := jwt.VerifyToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	err = helpers.DecodeBody(&body, context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	chosenProduct := product.SearchById(body.ID)
	if chosenProduct.Stock < body.Amount {
		context.JSON(http.StatusBadRequest, gin.H{"message": helpers.ProductNotEnoughStockError})
		context.Abort()
		return
	}

	cartDetails, err := c.CartService.UpdateProductInCart(body.ID, decodedToken.UserId, body.Amount)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	var newTotalPrice float64
	if cartDetails.Amount > body.Amount {
		difference := cartDetails.Amount - body.Amount
		newTotalPrice = cartDetails.TotalPrice - (float64(difference) * cartDetails.UnitPrice)
	} else {
		newTotalPrice = float64(body.Amount) * cartDetails.UnitPrice
	}

	cart.UpdateProductInCart(decodedToken.UserId, cartDetails.ProductId, body.Amount, newTotalPrice)

	// Update user main cart
	cart.UpdateUserCart(decodedToken.UserId, body.Amount, newTotalPrice)

	context.JSON(http.StatusOK, gin.H{"message": "Product updated in cart"})

}

// AddProductToCart adds product to cart
func (c *Controller) AddProductToCart(context *gin.Context) {
	var body RequestBody
	decodedToken, err := jwt.VerifyToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	err = helpers.DecodeBody(&body, context)
	if err != nil {
		return
	}

	chosenProduct, err := c.CartService.AddProductToCart(body.ID, body.Amount, decodedToken.UserId)

	if err != nil {
		if err == helpers.ProductAlreadyExistInCart {
			cartDetails, _ := c.CartService.UpdateProductInCart(body.ID, decodedToken.UserId, body.Amount)
			newAmount := body.Amount + cartDetails.Amount
			newTotalPrice := cartDetails.TotalPrice + (float64(newAmount) * cartDetails.UnitPrice)

			cart.UpdateProductInCart(decodedToken.UserId, cartDetails.ProductId, newAmount, newTotalPrice)
			cart.UpdateUserCart(decodedToken.UserId, body.Amount, newTotalPrice)

			context.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
			context.Abort()
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	newTotalPrice := chosenProduct.Price * float64(body.Amount)

	// Update user main cart
	fmt.Println(decodedToken)
	cart.UpdateUserCart(decodedToken.UserId, body.Amount, newTotalPrice)

	// Add product to cart details
	cart.CreateCartDetails(&cart.CartDetails{
		ProductName: chosenProduct.ProductName,
		CartId:      decodedToken.UserId,
		ProductId:   chosenProduct.ID,
		Amount:      body.Amount,
		UnitPrice:   chosenProduct.Price,
		TotalPrice:  newTotalPrice,
	})

	context.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})

}
