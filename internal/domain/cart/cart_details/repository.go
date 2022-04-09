package cartdetails

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/infrastructure"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var cardDetailsRepo *Repository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	cardDetailsRepo = NewRepository(db)
	cardDetailsRepo.Migration()
}

// NewRepository Creates CartDetails repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration Migrates CartDetails table
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&CartDetails{})
	if err != nil {
		panic(err)
	}
}

// NewModel Creates new CartDetails model
func NewModel(amount int, unitPrice float64, cartID uint, productID uint) *CartDetails {

	return &CartDetails{
		Amount:    amount,
		UnitPrice: unitPrice,
		ProductId: productID,
		CartId:    cartID,
	}
}

// Create creates new CartDetails model in database
func Create(model *CartDetails) {
	cardDetailsRepo.db.Create(model)
}

func FindCartDetails(cartID uint) []CartDetails {
	var cartDetails []CartDetails
	cardDetailsRepo.db.Where("cart_id = ?", cartID).Find(&cartDetails)

	return cartDetails
}

func IsProductExist(cartID uint, productID uint) bool {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	if cartDetails.ID == 0 {
		return false
	}

	return true
}

func UpdateProduct(cartID uint, productID uint, amount int) {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	cardDetailsRepo.db.Model(&cartDetails).Update("amount", amount)
}

func DeleteProduct(cartID uint, productID uint) {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	cardDetailsRepo.db.Delete(&cartDetails)
}
