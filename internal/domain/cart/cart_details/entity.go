package cartdetails

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/cart"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
	"gorm.io/gorm"
)

type CartDetails struct {
	gorm.Model
	Amount    int             `json:"amount"`
	UnitPrice float64         `json:"unit_price"`
	ProductId uint            `json:"product_id"`
	CartId    uint            `json:"cart_id"`
	Product   product.Product `gorm:"foreignkey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //
	Cart      cart.Cart       `gorm:"foreignkey:CartId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
