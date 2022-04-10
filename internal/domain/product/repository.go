package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/infrastructure"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var productRepo *Repository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	productRepo = NewRepository(db)
	productRepo.Migration()
}

// NewRepository Creates user repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration for product table
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
}

// NewModel Creates new product model
func NewModel(productName string, categoryName string, price float64, stock int, sku string) *Product {

	return &Product{
		ProductName:  productName,
		Price:        price,
		Stock:        stock,
		CategoryName: categoryName,
		SKU:          sku,
	}
}

// Create creates new product model in database
func Create(product *Product) {
	productRepo.db.Create(product)
}

// FindAll Finds all products in db
func FindAll() []Product {
	var products []Product
	productRepo.db.Find(&products)

	return products
}

// IsProductExist checks if product is already exist
func IsProductExist(productName string, sku string) bool {
	allProducts := FindAll()

	if len(allProducts) != 0 {
		var product Product
		productRepo.db.Where("product_name = ? OR sku = ?", productName, sku).Find(&product)
		if product.ID != 0 {
			return true
		}
	}
	return false
}

// SearchProduct searches products by product name and sku and returns []Product
func SearchProduct(queryString string) []Product {
	var products []Product
	productRepo.db.Where("product_name LIKE ?", "%"+queryString+"%").Or(productRepo.db.Where("sku LIKE ?", "%"+queryString+"%")).Find(&products)

	return products
}

// SearchById searches products by product id and returns Product
func SearchById(id uint) *Product {
	var product Product
	productRepo.db.Where("id = ?", id).Find(&product)

	return &product
}

// UpdateStock updates product stock
func UpdateStock(p Product, newStock int) {
	p.Stock = newStock
	productRepo.db.Save(&p)
}

// UpdateName updates product name
func UpdateName(p Product, newProductName string) {
	p.ProductName = newProductName
	productRepo.db.Save(&p)
}

// UpdatePrice updates product price
func UpdatePrice(p Product, newPrice float64) {
	p.Price = newPrice
	productRepo.db.Save(&p)
}

// UpdateSKU updates product sku
func UpdateSKU(p Product, newSKU string) {
	p.SKU = newSKU
	productRepo.db.Save(&p)
}

// UpdateCategory updates product category id
func UpdateCategory(p Product, newCategoryName string) {
	p.CategoryName = newCategoryName
	productRepo.db.Save(&p)
}

// DeleteProduct Deletes product from db
func DeleteProduct(p Product) {
	productRepo.db.Delete(&p)
}
