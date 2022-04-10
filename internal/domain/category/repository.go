package category

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/database_handler"
	"gorm.io/gorm"
	"strings"
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

// NewRepository Creates category repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration for category table
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		panic(err)
	}
}

// Create creates new category model in database
func Create(category *Category) {
	repository.db.Create(category)
}

// IsCategoryExist checks category is exist or not
func IsCategoryExist(categoryName string) bool {
	var model Category
	repository.db.Where("LOWER(category_name) = ?", strings.ToLower(categoryName)).Find(&model)

	if model.ID == 0 {
		return false
	}

	return true
}

// FindAll finds all category
func FindAll() []Category {
	var models []Category
	repository.db.Find(&models)

	return models
}
