package cart

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/user"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	TotalPrice float64   `json:"total_price" gorm:"default:0"`
	Amount     int       `json:"amount" gorm:"default:0"`
	UserId     uint      `json:"user_id"`
	User       user.User `gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
