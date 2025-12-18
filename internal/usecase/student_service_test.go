package usecase

import (
	"fmt"
	"testing"

	"github.com/alven-da/my-university-go/internal/adapter/repository"
	"github.com/alven-da/my-university-go/internal/domain"
)

type mockRepo struct {
  createdStudents []domain.Student
}

// mock the repository
func (m *mockRepo) Create(student domain.Student) (domain.Student, error) {
    student.ID = len(m.createdStudents) + 1
    m.createdStudents = append(m.createdStudents, student)
    return student, nil
}

func (m *mockRepo) GetByID(id int) (domain.Student, error) {
    for _, s := range m.createdStudents {
        if s.ID == id {
            return s, nil
        }
    }
    return domain.Student{}, fmt.Errorf("not found")
}

func TestRegisterStudent(t *testing.T) {
    repo := &mockRepo{}
    service := NewStudentService(repo)

    student, err := service.RegisterStudent("Alven", "test@example.com")

    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if student.ID != 1 {
        t.Errorf("expected ID 1, got %d", student.ID)
    }

    if student.Name != "Alven" {
        t.Errorf("expected name 'Alven', got %s", student.Name)
    }

    if student.Email != "test@example.com" {
        t.Errorf("expected email 'test@example.com', got %s", student.Email)
    }
}

func TestGetStudentByID(t *testing.T) {
	repo := repository.NewMemoryRepoWithMockData()

	service := NewStudentService(repo)

	student, err := service.GetStudentById(2)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if student.Name != "Sheryl" {
		t.Errorf("expected name 'Sheryl', got %s", student.Name)
	}

	if student.Email != "sheryl@test.com" {
		t.Errorf("expected email 'sheryl@test.com', got %s", student.Email)
	}
}