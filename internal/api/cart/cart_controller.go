package cart

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	CartService *cart.Service
}

func NewCartController(cartService *cart.Service) *Controller {
	return &Controller{
		CartService: cartService,
	}
}

func (c *Controller) GetCartList(context *gin.Context) {
	err := c.CartService.GetCartList()
	if err != nil {
		context.JSON(http.StatusNoContent, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

}

func (c *Controller) DeleteItemFromCart(context *gin.Context) {

}

func (c *Controller) UpdateItemInCart(context *gin.Context) {

}

func (c *Controller) AddItemToCart(context *gin.Context) {

}
