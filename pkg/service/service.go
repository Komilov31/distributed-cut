package service

type Cut interface {
	ProcessFlagF(nextLine string) string
}

type Service struct {
	cut Cut
}

func New(cut Cut) *Service {
	return &Service{
		cut: cut,
	}
}

func (s *Service) ProcessInput(input map[int]string) map[int]string {
	result := make(map[int]string, len(input))

	for key, value := range input {
		processed := s.cut.ProcessFlagF(value)
		result[key] = processed
	}

	return result
}
