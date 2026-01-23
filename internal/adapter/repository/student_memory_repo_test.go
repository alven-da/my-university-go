package repository

import (
	"testing"

	"github.com/alven-da/my-university-go/internal/domain"
)

func TestMemoryRepo_CreateAndGetByID(t *testing.T) {
    repo := NewMemoryRepo()

    created, err := repo.Create(domain.Student{Name: "Alven", Email: "alven@test.com"})
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if created.ID != 1 {
        t.Errorf("expected ID 1, got %d", created.ID)
    }

    fetched, err := repo.GetByID(created.ID)
    if err != nil {
        t.Fatalf("unexpected error fetching student: %v", err)
    }

    if fetched.Name != "Alven" {
        t.Errorf("expected name 'Alven', got %s", fetched.Name)
    }
    if fetched.Email != "alven@test.com" {
        t.Errorf("expected email 'alven@test.com', got %s", fetched.Email)
    }
}

func TestMemoryRepo_GetByID_NotFound(t *testing.T) {
    repo := NewMemoryRepo()

    _, err := repo.GetByID(999)
    if err == nil {
        t.Errorf("expected error for non-existing student, got nil")
    }
}
