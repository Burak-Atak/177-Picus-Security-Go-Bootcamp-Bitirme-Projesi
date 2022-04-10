package category

import "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateCategory(categoryName string) (*Category, error) {

	if IsCategoryExist(categoryName) {
		return nil, helpers.CategoryAlreadyExistError
	}

	category := GetCategoryByName(categoryName)
	return category, nil
}

func (s *Service) GetCategoryByName(categoryName string) (*Category, error) {

	if !IsCategoryExist(categoryName) {
		return nil, helpers.CategoryNotFoundError
	}

	category := GetCategoryByName(categoryName)
	return category, nil
}
