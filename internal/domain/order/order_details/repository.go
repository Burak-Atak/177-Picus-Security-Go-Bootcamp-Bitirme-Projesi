package orderdetails

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

// Migration for order details table
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&OrderDetails{})
	if err != nil {
		return
	}
}

// NewModel Creates new order model and adds it to database
func NewModel(orderId uint, productId uint, amount int, unitPrice float64, totalPrice float64, productName string) {
	newModel := &OrderDetails{
		OrderId:     orderId,
		ProductId:   productId,
		Amount:      amount,
		UnitPrice:   unitPrice,
		TotalPrice:  totalPrice,
		ProductName: productName,
	}

	repository.db.Create(newModel)
}

// FindOrderDetails finds order details by order id and returns OrderDetails
func FindOrderDetails(orderId uint) []OrderDetails {
	var orderDetails []OrderDetails
	repository.db.Where("order_id = ?", orderId).Find(&orderDetails)

	return orderDetails
}

// DeleteModel deletes OrderDetails model from database
func DeleteModel(orderDetail OrderDetails) {
	repository.db.Delete(&orderDetail)
}
