package cart

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
	"gorm.io/gorm"
)

type CartDetails struct {
	gorm.Model
	ProductName string          `json:"product_name"`
	Amount      int             `json:"amount"`
	UnitPrice   float64         `json:"unit_price"`
	TotalPrice  float64         `json:"total_price"`
	ProductId   uint            `json:"product_id"`
	CartId      uint            `json:"cart_id"`
	Product     product.Product `gorm:"foreignkey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //
	Cart        Cart            `gorm:"foreignkey:CartId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
