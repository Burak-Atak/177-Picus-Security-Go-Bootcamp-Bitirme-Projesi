package product

import (
	"fmt"

	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/abc/infrastructure"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

var productRepo *ProductRepository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	productRepo = NewRepository(db)
	productRepo.Migration()
}

// Creates user repository
func NewRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// Migration
func (r *ProductRepository) Migration() {
	r.db.AutoMigrate(&Product{})
}

// Creates new product model and adds it to database
func NewModel(productName string, categoryName string, price float64, stock int, sku string) {

	newProduct := &Product{
		ProductName:  productName,
		Price:        price,
		Stock:        stock,
		CategoryName: categoryName,
		SKU:          sku,
	}

	productRepo.db.Create(newProduct)
}

// Finds all products in db
func FindAll() []Product {
	var products []Product
	productRepo.db.Find(&products)

	return products
}

// IsProductExist checks if product is already exist
func IsProductExist(productName string) bool {
	allProducts := FindAll()

	if len(allProducts) != 0 {
		var product Product
		productRepo.db.Where("product_name = ?", productName).Find(&product)
		fmt.Println(product)
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

// SearchProductByID searches products by product id and returns Product
func SearchProductByID(id uint) Product {
	var product Product
	productRepo.db.Where("id = ?", id).Find(&product)

	return product
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

// UpdateCategoryID updates product category id
func UpdateCategoryID(p Product, newCategoryName string) {
	p.CategoryName = newCategoryName
	productRepo.db.Save(&p)
}

// Deletes product from db
func DeleteProduct(p Product) {
	productRepo.db.Delete(&p)
}
