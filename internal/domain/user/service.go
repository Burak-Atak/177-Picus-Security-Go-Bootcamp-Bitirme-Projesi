package user

import "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// CreateUser checks if the user exists
func (s *Service) CreateUser(user *User) error {
	if IsUserExist(user.Email) {
		return helpers.UserExistsError
	}

	return nil
}

// Login checks if the user exists and if the password is correct
func (s *Service) Login(email string, password string) (*User, error) {
	user := SearchByEmail(email)
	if user.ID == 0 {
		return nil, helpers.UserNotFoundError
	}

	if user.Password != password {
		return nil, helpers.InvalidPasswordError
	}
	return &user, nil
}
