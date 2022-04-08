package cartdetails

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/abc/domain/cart"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/abc/domain/product"
)

type CartDetails struct {
	ID        uint            `gorm:"primary_key"`
	ProductID uint            `json:"product_id"`
	Amount    int             `json:"amount"`
	CartID    uint            `json:"cart_id"`
	Product   product.Product `gorm:"foreignkey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cart      cart.Cart       `gorm:"foreignkey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
