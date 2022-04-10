package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/category"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName  string            `json:"product_name" validate:"required"`
	Price        float64           `json:"price" validate:"gt=1"`
	Stock        int               `json:"stock" validate:"gt=1"`
	CategoryName string            `json:"category_name" validate:"required"`
	SKU          string            `json:"sku" validate:"required"`
	Category     category.Category `gorm:"foreignkey:CategoryName;references:CategoryName; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
