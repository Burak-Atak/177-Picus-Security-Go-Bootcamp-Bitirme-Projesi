package order

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/database_handler"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var repository *Repository

func init() {
	db := database_handler.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	repository = NewRepository(db)
	repository.Migration()
}

// NewRepository Creates order repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

//Migration for order table
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}
}

// NewModel Creates new order model
func NewModel(totalPrice float64, amount int, userId uint) *Order {
	return &Order{
		TotalPrice: totalPrice,
		Amount:     amount,
		UserId:     userId,
	}
}

// Create new order model in database
func Create(order *Order) {
	repository.db.Create(order)
}

// SearchById Searches order by id
func SearchById(id uint, userId uint) *Order {
	var model Order
	repository.db.Where("id = ? AND user_id = ?", id, userId).Find(&model)

	return &model
}

// DeleteOrder deletes order by id
func DeleteOrder(id uint) {
	var model Order
	repository.db.Where("id = ?", id).Delete(&model)
}

// FindUserOrders finds all orders of user
func FindUserOrders(userId uint) []Order {
	var orders []Order
	repository.db.Where("user_id = ?", userId).Find(&orders)

	return orders
}
