package usecase

import (
	"fmt"

	"github.com/alven-da/my-university-go/internal/domain"
)

type SubjectService struct {
	repo domain.SubjectRepository
}

func NewSubjectService(r domain.SubjectRepository) *SubjectService {
	return &SubjectService{repo: r}
}

func (s *SubjectService) GetSubjectById(id int) (domain.Subject, error)  {
	data, err := s.repo.GetByID(id)

	if err != nil {
		return domain.Subject{}, fmt.Errorf("Failed to get subject with id %s: %w", id, err)
	}

	return data, nil
}