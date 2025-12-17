package repository

import (
	"fmt"

	"github.com/alven-da/my-university-go/internal/domain"
)

type MemoryRepo struct {
    users map[int]domain.User
    nextID int
}

func NewMemoryRepo() *MemoryRepo {
    return &MemoryRepo{users: make(map[int]domain.User), nextID: 1}
}

func (r *MemoryRepo) Create(user domain.User) (domain.User, error) {
    user.ID = r.nextID
    r.users[r.nextID] = user
    r.nextID++
    return user, nil
}

func (r *MemoryRepo) GetByID(id int) (domain.User, error) {
    user, ok := r.users[id]
    if !ok {
        return domain.User{}, fmt.Errorf("not found")
    }
    return user, nil
}