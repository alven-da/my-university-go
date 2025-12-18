package repository

import (
	"fmt"

	"github.com/alven-da/my-university-go/internal/domain"
)

type MemoryRepo struct {
    students map[int]domain.Student
    nextID int
}

func NewMemoryRepo() *MemoryRepo {
    return &MemoryRepo{students: make(map[int]domain.Student), nextID: 1}
}

func NewMemoryRepoWithMockData() *MemoryRepo {
    repo := &MemoryRepo{students: make(map[int]domain.Student), nextID: 1}

    mockStudents := []domain.Student{
        { ID: 1, Name: "Alven", Email: "alven@test.com" },
        { ID: 1, Name: "Sheryl", Email: "sheryl@test.com" },
    }

    for _, s := range mockStudents {
        repo.Create(s)
    }

    return repo
}

// extends the methods
func (r *MemoryRepo) Create(student domain.Student) (domain.Student, error) {
    student.ID = r.nextID
    r.students[r.nextID] = student
    r.nextID++
    return student, nil
}

func (r *MemoryRepo) GetByID(id int) (domain.Student, error) {
    student, ok := r.students[id]
    if !ok {
        return domain.Student{}, fmt.Errorf("not found")
    }
    return student, nil
}