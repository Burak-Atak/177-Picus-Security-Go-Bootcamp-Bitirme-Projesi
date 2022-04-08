package cart

import "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/abc/domain/user"

type Cart struct {
	ID         uint      `gorm:"primary_key"`
	TotalPrice float64   `json:"total_price" gorm:"default:0"`
	UserID     uint      `json:"user_id"`
	User       user.User `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
