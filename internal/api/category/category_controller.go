package category

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/category"
)

type Controller struct {
	CategoryService *category.Service
}

func NewCategoryController(categoryService *category.Service) *Controller {
	return &Controller{
		CategoryService: categoryService,
	}
}
