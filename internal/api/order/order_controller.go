package order

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/order"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	OrderService *order.Service
}

func NewOrderController(orderService *order.Service) *Controller {
	return &Controller{
		OrderService: orderService,
	}
}

func (c *Controller) GetOrderList(context *gin.Context) {

}

func (c *Controller) GetOrderDetail(context *gin.Context) {

}

func (c *Controller) CancelOrder(context *gin.Context) {

}

func (c *Controller) CreateOrder(context *gin.Context) {

}
