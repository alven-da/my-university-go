package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alven-da/my-university-go/internal/domain"
	"github.com/alven-da/my-university-go/internal/usecase"
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

// Test cases
func TestRegisterStudentHandler(t *testing.T) {
	mockRepo := &mockRepo{}
	service := usecase.NewStudentService(mockRepo)

	handler := NewStudentHandler(service)

  body := []byte(`{"name":"Alven","email":"alven@test.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewBuffer(body))
  w := httptest.NewRecorder()

	handler.Register(w, req)

	res := w.Result()
  defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	var student domain.Student
	if err := json.NewDecoder(res.Body).Decode(&student); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if student.Name != "Alven" {
		t.Errorf("expected name 'Alven', got %s", student.Name)
	}
	if student.Email != "alven@test.com" {
		t.Errorf("expected email 'alven@test.com', got %s", student.Email)
	}
}

func TestGetByIdHandler(t *testing.T) {
	mockRepo := &mockRepo{}
	service := usecase.NewStudentService(mockRepo)
	
	handler := NewStudentHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/student/{id}", nil)
	w := httptest.NewRecorder()

	handler.Register(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}
}