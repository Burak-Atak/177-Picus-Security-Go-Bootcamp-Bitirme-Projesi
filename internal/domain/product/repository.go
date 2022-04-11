package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/database_handler"
	"gorm.io/gorm"
	"strings"
)

type Repository struct {
	db *gorm.DB
}

var productRepo *Repository

func init() {
	db := database_handler.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
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

// Update updates product model in database
func Update(product *Product) {
	productRepo.db.Save(product)
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
		productRepo.db.Where("LOWER(product_name) = ? OR LOWER(sku) = ?", strings.ToLower(productName), strings.ToLower(sku)).Find(&product)
		if product.ID != 0 {
			return true
		}
	}
	return false
}

func GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product

	allProducts := FindAll()
	productRepo.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products)

	return products, len(allProducts)
}

// SearchProduct searches products by product name and sku and returns []Product
func SearchProduct(queryString string) []Product {
	var products []Product
	productRepo.db.Where("LOWER(product_name) LIKE ?", "%"+strings.ToLower(queryString)+"%").Or(productRepo.db.Where("LOWER(sku) LIKE ?", "%"+strings.ToLower(queryString)+"%")).Find(&products)

	return products
}

// SearchProductWithPagination searches products by product name and sku and returns []Product
func SearchProductWithPagination(queryString string, pageIndex, pageSize int) []Product {
	var products []Product
	productRepo.db.Where("LOWER(product_name) LIKE ?", "%"+strings.ToLower(queryString)+"%").Or(productRepo.db.Where("LOWER(sku) LIKE ?", "%"+strings.ToLower(queryString)+"%")).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products)

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

// DeleteProduct Deletes product from db
func DeleteProduct(p Product) {
	productRepo.db.Delete(&p)
}

// SearchBySKU searches product by sku and returns Product
func SearchBySKU(sku string) *Product {
	var product Product
	productRepo.db.Where("LOWER(sku) = ?", strings.ToLower(sku)).Find(&product)

	return &product
}

// SearchByProductName searches products by product name and returns Product
func SearchByProductName(productName string) *Product {
	var product Product
	productRepo.db.Where("LOWER(product_name) = ?", strings.ToLower(productName)).Find(&product)

	return &product
}
