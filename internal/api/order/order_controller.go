package order

import (
	"fmt"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/cart"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/order"
	orderdetails "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/order/order_details"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	OrderService *order.Service
}

func NewOrderController(orderService *order.Service) *Controller {
	return &Controller{
		OrderService: orderService,
	}
}

// GetOrderList gets user's order list
func (c *Controller) GetOrderList(context *gin.Context) {
	decodedToken, _ := jwt.VerifyToken(context.GetHeader("Authorization"))

	userOrders := order.FindUserOrders(decodedToken.UserId)
	if len(userOrders) == 0 {
		context.JSON(http.StatusOK, gin.H{
			"message": helpers.OrderNotFoundError,
		})
		context.Abort()
		return
	}

	outPut := make([]Order, len(userOrders))
	for i, chosenOrder := range userOrders {
		outPut[i] = Order{
			TotalPrice: chosenOrder.TotalPrice,
			Amount:     chosenOrder.Amount,
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"Your orders": outPut,
	})
}

// GetOrderDetails is used to get order details
func (c *Controller) GetOrderDetails(context *gin.Context) {
	decodedToken, _ := jwt.VerifyToken(context.GetHeader("Authorization"))
	id, isOk := context.GetQuery("id")
	if !isOk {
		context.JSON(http.StatusBadRequest, helpers.IdIsRequiredError.Error())
		context.Abort()
		return
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.InvalidIdError.Error())
		context.Abort()
		return
	}

	_, err = c.OrderService.GetOrder(uint(orderId), decodedToken.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		context.Abort()
		return
	}
	orderDetails := orderdetails.FindOrderDetails(uint(orderId))
	outPut := make([]Product, len(orderDetails))
	for i, orderDetail := range orderDetails {
		outPut[i] = Product{
			ProductName: orderDetail.ProductName,
			Amount:      orderDetail.Amount,
			UnitPrice:   orderDetail.UnitPrice,
			TotalPrice:  orderDetail.TotalPrice,
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"Order Details": outPut,
	})

}

// CancelOrder cancels the order
func (c *Controller) CancelOrder(context *gin.Context) {
	decodedToken, _ := jwt.VerifyToken(context.GetHeader("Authorization"))
	id, isOk := context.GetQuery("id")
	if !isOk {
		context.JSON(http.StatusBadRequest, helpers.IdIsRequiredError.Error())
		context.Abort()
		return
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.InvalidIdError.Error())
		context.Abort()
		return
	}

	err = c.OrderService.CancelOrder(uint(orderId), decodedToken.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		context.Abort()
		return
	}
	// Deletes order details, updates product stock
	chosenOrderDetails := orderdetails.FindOrderDetails(uint(orderId))
	for _, chosenOrderDetail := range chosenOrderDetails {
		chosenProduct := product.SearchById(chosenOrderDetail.ProductId)
		chosenProduct.Stock += chosenOrderDetail.Amount
		product.Update(chosenProduct)
		orderdetails.DeleteModel(chosenOrderDetail)
	}

	// Deletes order
	order.DeleteOrder(uint(orderId))

	context.JSON(http.StatusOK, gin.H{
		"message": "Order has been canceled",
	})
}

// CreateOrder is a function that creates an order
func (c *Controller) CreateOrder(context *gin.Context) {
	cartService := cart.NewService()
	productService := product.NewService()
	decodedToken, _ := jwt.VerifyToken(context.GetHeader("Authorization"))

	allProductsInCart, err := cartService.GetCartsHasCartId(decodedToken.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	// Check if products in cart are available
	var outOfStockProducts []string
	for _, productInCart := range *allProductsInCart {
		chosenProduct, _ := productService.GetProduct(productInCart.ProductId)
		if chosenProduct.Stock < productInCart.Amount {
			newMessage := fmt.Sprintf("Product %s has no enought stock. Available stock is %d, your amount is %d",
				chosenProduct.ProductName, chosenProduct.Stock, productInCart.Amount)
			outOfStockProducts = append(outOfStockProducts, newMessage)
		}
	}
	if len(outOfStockProducts) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message":  helpers.NotEnoughStockError.Error(),
			"products": outOfStockProducts,
		})
		context.Abort()
		return
	}

	userCart := cart.SearchById(decodedToken.UserId)
	newOrder := order.NewModel(userCart.TotalPrice, userCart.Amount, userCart.UserId)
	order.Create(newOrder)
	userCart.TotalPrice = 0
	userCart.Amount = 0
	cart.Update(userCart)

	// Creates order details, updates product stock, and delete products from cart
	for _, productInCart := range *allProductsInCart {
		chosenProduct, _ := productService.GetProduct(productInCart.ProductId)
		newStock := chosenProduct.Stock - productInCart.Amount
		orderdetails.NewModel(newOrder.ID, chosenProduct.ID, productInCart.Amount, chosenProduct.Price, productInCart.TotalPrice, chosenProduct.ProductName)
		product.UpdateStock(*chosenProduct, newStock)
		cart.DeleteModel(&productInCart)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
	})

}
