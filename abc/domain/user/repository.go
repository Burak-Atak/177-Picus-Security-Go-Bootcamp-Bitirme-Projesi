package user

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/abc/infrastructure"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var userRepo *UserRepository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	userRepo = NewRepository(db)
	userRepo.Migration()
}

// Creates user repository
func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Migration
func (r *UserRepository) Migration() {
	r.db.AutoMigrate(&User{})
}

// Creates new user model and adds it to database
func NewModel(email string, password string, role string) {
	newUser := &User{
		Email:    email,
		Password: password,
		Role:     role,
	}

	userRepo.db.Create(newUser)
}

func SearchById(id uint) User {
	var user User
	userRepo.db.Where("id = ?", id).First(&user)

	return user
}
