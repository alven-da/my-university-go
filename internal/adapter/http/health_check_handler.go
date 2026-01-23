package http

import (
	"encoding/json"
	"net/http"

	"github.com/alven-da/my-university-go/internal/usecase"
)

type PublicHandler struct {
	service *usecase.PublicService
}

func NewPublicHandler(s *usecase.PublicService) *PublicHandler {
	return &PublicHandler{service: s}
}

func (h *PublicHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := h.service.HealthCheck()

	json.NewEncoder(w).Encode(data)
}