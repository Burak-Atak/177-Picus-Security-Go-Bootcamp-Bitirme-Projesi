package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=20"`
	Role     string `json:"role" validate:"oneof=admin costumer"`
}
