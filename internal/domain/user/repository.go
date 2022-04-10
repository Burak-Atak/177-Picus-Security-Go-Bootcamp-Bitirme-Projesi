package user

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/infrastructure"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

var repository *Repository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	repository = NewRepository(db)
	repository.Migration()
}

// NewRepository Creates user repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration for user table
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

// NewModel Creates new user model
func NewModel(email string, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

// Create new user model to database
func Create(user *User) {
	repository.db.Create(user)
}

// SearchById search user by id
func SearchById(id uint) User {
	var model User
	repository.db.Where("id = ?", id).Find(&model)

	return model
}

// IsUserExist checks if user exist
func IsUserExist(email string) bool {
	var model User
	repository.db.Where("email = ?", email).Find(&model)

	if model.ID == 0 {
		return false
	}

	return true
}

func SearchByEmail(email string) User {
	var model User
	repository.db.Where("email = ?", email).Find(&model)

	return model
}
