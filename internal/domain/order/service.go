package order

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// GetOrder checks if order exists and returns it
func (s *Service) GetOrder(orderId uint, userId uint) (*Order, error) {
	order := SearchById(orderId, userId)
	if order.ID == 0 {
		return nil, helpers.OrderNotFoundError
	}

	return order, nil
}

// CancelOrder checks if date is not passed and order exists
func (s *Service) CancelOrder(orderId uint, userId uint) error {
	order := SearchById(orderId, userId)
	if order.ID == 0 {
		return helpers.OrderNotFoundError
	}
	currentTime := time.Now()
	timeDifference := currentTime.Sub(currentTime).Hours()

	if timeDifference > 24*14 {
		return helpers.OrderCancelError

	}

	return nil
}
