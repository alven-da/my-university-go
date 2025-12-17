package http

import (
	"encoding/json"
	"net/http"

	"github.com/alven-da/my-university-go/internal/usecase"
	"github.com/gorilla/mux"
)

type StudentHandler struct {
    service *usecase.StudentService
}

func NewStudentHandler(s *usecase.StudentService) *StudentHandler {
    return &StudentHandler{service: s}
}

func (h *StudentHandler) Register(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var input struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    _ = json.NewDecoder(r.Body).Decode(&input)

    student, err := h.service.RegisterUser(input.Name, input.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(student)
}

func (h *StudentHandler) GetById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    vars := mux.Vars(r)

    studentId := vars["id"]

    data := struct {
        Id string `json:"id"`
    }{
        Id: studentId,
    }

    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}