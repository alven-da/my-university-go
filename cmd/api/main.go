package main

import (
	"net/http"

	user_handler "github.com/alven-da/my-university-go/internal/adapter/http"
	"github.com/alven-da/my-university-go/internal/adapter/repository"
	"github.com/alven-da/my-university-go/internal/usecase"
)

func main() {
    repo := repository.NewMemoryRepo()
    service := usecase.NewUserService(repo)
    handler := user_handler.NewUserHandler(service)

    http.HandleFunc("/users", handler.Register)
    http.ListenAndServe(":8080", nil)
}