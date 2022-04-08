package product

type Product struct {
	ID           uint    `gorm:"primary_key"`
	ProductName  string  `json:"product_name"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryName string  `json:"category_name"`
	SKU          string  `json:"sku"`
}
