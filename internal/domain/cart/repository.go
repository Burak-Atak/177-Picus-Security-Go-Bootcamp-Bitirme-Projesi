package cart

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/infrastructure"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var cartRepo *Repository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	cartRepo = NewRepository(db)
	cartRepo.Migration()
}

// NewCartDetailsRepository Creates cart repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration Migrates cart table
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Cart{})
	if err != nil {
		return
	}
}

// NewCartDetailsModel Creates new cart model
func NewModel(userId uint) *Cart {

	return &Cart{
		UserId: userId,
	}
}

// CreateCartDetails creates new cart model in database
func Create(cart *Cart) {
	cartRepo.db.Create(cart)
}

// SearchById searches cart by id
func SearchById(id uint) *Cart {
	var cart Cart
	cartRepo.db.Where("id = ?", id).First(&cart)

	return &cart
}

// IsCartExist checks if cart exists
func IsCartExist(userId uint) bool {
	var cart Cart
	cartRepo.db.Where("user_id = ?", userId).First(&cart)

	if cart.ID == 0 {
		return false
	}

	return true
}

func Update(cart *Cart) {
	cartRepo.db.Save(cart)
}

func UpdateUserCart(userId uint, newAmount int, newPrice float64) {
	usersCart := SearchById(userId)
	usersCart.Amount += newAmount
	usersCart.TotalPrice += newPrice
	Update(usersCart)
}
