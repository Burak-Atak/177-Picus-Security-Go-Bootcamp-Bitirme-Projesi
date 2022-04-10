package category

import "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// GetCategoryByName checks if category exists
func (s *Service) GetCategoryByName(categoryName string) error {

	if IsCategoryExist(categoryName) {
		return helpers.CategoryAlreadyExistError
	}
	return nil
}
