package repository

import (
	"fmt"

	"github.com/alven-da/my-university-go/internal/domain"
)

type SubjectMemoryRepo struct {
	subjects map[int]domain.Subject
	nextID int
}

func NewSubjectMemoryRepo() *SubjectMemoryRepo {
	return &SubjectMemoryRepo{
		subjects: make(map[int]domain.Subject),
		nextID: 1,
	}
}

func NewSubjectMemoryRepoWithMockData() *SubjectMemoryRepo {
    repo := &SubjectMemoryRepo{subjects: make(map[int]domain.Subject), nextID: 1}

    mockStudents := []domain.Subject{
        { ID: 1, Name: "English 101" },
        { ID: 2, Name: "Literature 101" },
    }

    for _, s := range mockStudents {
        repo.Create(s)
    }

    return repo
}

// extend methods
func (r *SubjectMemoryRepo) Create(s domain.Subject) (domain.Subject, error) {
	s.ID = r.nextID
	r.subjects[r.nextID] = s
	r.nextID++

	return s, nil
}

func (r *SubjectMemoryRepo) GetByID(id int) (domain.Subject, error) {
	subject, ok := r.subjects[id]

	if (!ok) {
		return domain.Subject{}, fmt.Errorf("not found")
	}

	return subject, nil
}