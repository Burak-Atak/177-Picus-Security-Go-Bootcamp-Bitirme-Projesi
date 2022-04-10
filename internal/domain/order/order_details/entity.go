package orderdetails

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/order"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
)

type OrderDetails struct {
	ID          uint            `gorm:"primary_key"`
	OrderId     uint            `json:"order_id"`
	ProductId   uint            `json:"product_id"`
	ProductName string          `json:"product_name"`
	Amount      int             `json:"amount"`
	UnitPrice   float64         `json:"unit_price"`
	TotalPrice  float64         `json:"total_price"`
	Order       order.Order     `gorm:"foreignkey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product     product.Product `gorm:"foreignkey:ProductId"`
}
