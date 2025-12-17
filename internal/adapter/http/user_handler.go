package http

import (
	"encoding/json"
	"net/http"

	"github.com/alven-da/my-university-go/internal/usecase"
)

type UserHandler struct {
    service *usecase.UserService
}

func NewUserHandler(s *usecase.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    _ = json.NewDecoder(r.Body).Decode(&input)

    user, err := h.service.RegisterUser(input.Name, input.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}