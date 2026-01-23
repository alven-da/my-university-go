package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alven-da/my-university-go/internal/usecase"
	"github.com/gorilla/mux"
)

type SubjectHandler struct {
	service *usecase.SubjectService
}

func NewSubjectHandler(s *usecase.SubjectService) *SubjectHandler {
	return &SubjectHandler{service: s}
}

func (h *SubjectHandler) GetSubjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	subject_id := vars["id"]

	subject_id_int, err := strconv.Atoi(subject_id)

	if err != nil {
		http.Error(w, "Invalid subject ID", http.StatusBadRequest)
		return
	}

	subject, err := h.service.GetSubjectById(subject_id_int)

	if err != nil {
		fmt.Errorf("Error in fetching subject data with id %s: %w", subject_id, err)
	}

	dto := struct {
		SubjectID int `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
	}{
		SubjectID: subject.ID,
		Name: subject.Name,
		Description: subject.Description,
	}

	if err := json.NewEncoder(w).Encode(dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}