package usecase

import (
	"fmt"

	"github.com/alven-da/my-university-go/internal/domain"
)

type StudentService struct {
    repo domain.StudentRepository
}

func NewStudentService(r domain.StudentRepository) *StudentService {
    return &StudentService{repo: r}
}

func (s *StudentService) RegisterUser(name, email string) (domain.Student, error) {
    user := domain.Student{Name: name, Email: email}
    return s.repo.Create(user)
}

func (s *StudentService) GetStudentById(id int) (domain.Student, error) {
    student, err := s.repo.GetByID(id)

    if err != nil {
        return domain.Student{}, fmt.Errorf("failed to get student with id %d: %w", id, err)
    }

    return student, nil
}