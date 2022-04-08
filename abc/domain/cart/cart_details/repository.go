package cartdetails

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/abc/infrastructure"
	"gorm.io/gorm"
)

type CartDetailsRepository struct {
	db *gorm.DB
}

var cardDetailsRepo *CartDetailsRepository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	cardDetailsRepo = NewRepository(db)
	cardDetailsRepo.Migration()
}

// Creates CartDetails repository
func NewRepository(db *gorm.DB) *CartDetailsRepository {
	return &CartDetailsRepository{
		db: db,
	}
}

// Migration
func (r *CartDetailsRepository) Migration() {
	r.db.AutoMigrate(&CartDetails{})
}

// Creates new CartDetails model and adds it to database
func NewModel(amount int, unitPrice float64, cartID uint, productID uint) {

	newCardDetails := &CartDetails{
		Amount:    amount,
		ProductID: productID,
		CartID:    cartID,
	}

	cardDetailsRepo.db.Create(newCardDetails)
}
