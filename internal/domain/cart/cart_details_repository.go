package cart

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/infrastructure"
	"gorm.io/gorm"
)

type CartDetailsRepository struct {
	db *gorm.DB
}

var cardDetailsRepo *CartDetailsRepository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	cardDetailsRepo = NewCartDetailsRepository(db)
	cardDetailsRepo.Migration()
}

// NewCartDetailsRepository Creates CartDetails repository
func NewCartDetailsRepository(db *gorm.DB) *CartDetailsRepository {
	return &CartDetailsRepository{
		db: db,
	}
}

// Migration Migrates CartDetails table
func (r *CartDetailsRepository) Migration() {
	err := r.db.AutoMigrate(&CartDetails{})
	if err != nil {
		panic(err)
	}
}

// NewCartDetailsModel Creates new CartDetails model
func NewCartDetailsModel(amount int, unitPrice float64, totalPrice float64, cartID uint, productID uint) *CartDetails {

	return &CartDetails{
		Amount:     amount,
		UnitPrice:  unitPrice,
		TotalPrice: totalPrice,
		ProductId:  productID,
		CartId:     cartID,
	}
}

// CreateCartDetails creates new CartDetails model in database
func CreateCartDetails(model *CartDetails) {
	cardDetailsRepo.db.Create(model)
}

// IsProductExist checks if product is already in cart
func IsProductExist(userId uint, productID uint) bool {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", userId, productID).Find(&cartDetails)

	if cartDetails.ID == 0 {
		return false
	}

	return true
}

// DeleteProductInCart deletes CartDetails model in database
func DeleteProductInCart(cartID uint, productID uint) {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	cardDetailsRepo.db.Delete(&cartDetails)
}

// GetAllCartDetailsOfUser gets CartDetails models by cartID
func GetAllCartDetailsOfUser(cartID uint) []CartDetails {
	var cartDetails []CartDetails
	cardDetailsRepo.db.Where("cart_id = ?", cartID).Find(&cartDetails)

	return cartDetails
}

func GetCartDetailsByCartIdAndProductId(cartID uint, productID uint) *CartDetails {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	return &cartDetails
}

// UpdateProductInCart updates CartDetails model in database
func UpdateProductInCart(cartID uint, productID uint, amount int, totalPrice float64) {
	var cartDetails CartDetails
	cardDetailsRepo.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Find(&cartDetails)

	cardDetailsRepo.db.Model(&cartDetails).Update("amount", amount)
	cardDetailsRepo.db.Model(&cartDetails).Update("total_price", totalPrice)
}

func GetCartDetailsByProductId(productID uint) *[]CartDetails {
	var cartDetails []CartDetails
	cardDetailsRepo.db.Where("product_id = ?", productID).Find(&cartDetails)

	return &cartDetails
}
