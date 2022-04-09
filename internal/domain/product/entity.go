package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/category"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string            `json:"product_name"`
	Price       float64           `json:"price"`
	Stock       int               `json:"stock"`
	CategoryId  uint              `json:"category_id"`
	SKU         string            `json:"sku"`
	Category    category.Category `gorm:"foreignkey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
