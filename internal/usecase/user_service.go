// internal/usecase/user_service.go
package usecase

import "github.com/alven-da/my-university-go/internal/domain"

type UserService struct {
    repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) RegisterUser(name, email string) (domain.User, error) {
    user := domain.User{Name: name, Email: email}
    return s.repo.Create(user)
}