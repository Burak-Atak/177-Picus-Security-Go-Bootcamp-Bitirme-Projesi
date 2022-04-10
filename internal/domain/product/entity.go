package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/category"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName  string            `json:"product_name" validate:"required"`
	Price        float64           `json:"price" validate:"required, number"`
	Stock        int               `json:"stock" validate:"required, number"`
	CategoryName string            `json:"category_id" validate:"required"`
	SKU          string            `json:"sku" validate:"required"`
	Category     category.Category `gorm:"references:CategoryName; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
