package order

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/user"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TotalPrice float64   `json:"total_price"`
	Amount     int       `json:"amount"`
	UserId     uint      `json:"user_id"`
	User       user.User `json:"user" gorm:"foreignkey:UserId"`
}
