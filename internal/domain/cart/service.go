package cart

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetCartList() error {
	return nil
}
