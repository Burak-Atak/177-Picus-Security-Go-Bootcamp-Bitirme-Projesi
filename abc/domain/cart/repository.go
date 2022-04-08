package cart

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/abc/infrastructure"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

var cartRepo *CartRepository

func init() {
	db := infrastructure.NewMySqlDB("root:mysql@tcp(127.0.0.1:3306)/application?charset=utf8mb4&parseTime=True&loc=Local")
	cartRepo = NewRepository(db)
	cartRepo.Migration()
}

// Creates cart repository
func NewRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

// Migration
func (r *CartRepository) Migration() {
	r.db.AutoMigrate(&Cart{})
}

func NewModel() {

	newCart := &Cart{}

	cartRepo.db.Create(newCart)
}
