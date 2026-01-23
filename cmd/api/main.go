package main

import (
	"log"
	"net/http"

	http_handler "github.com/alven-da/my-university-go/internal/adapter/http"
	"github.com/alven-da/my-university-go/internal/adapter/repository"
	"github.com/alven-da/my-university-go/internal/usecase"
	"github.com/gorilla/mux"
)

func main() {
    // initialize repo
    repo := repository.NewMemoryRepo()
    
    // initialize services
    student_service := usecase.NewStudentService(repo)
    public_service := usecase.NewPublicService()

    // initialize http handlers
    student_handler := http_handler.NewStudentHandler(student_service)
    public_handler := http_handler.NewPublicHandler(public_service)

    r := mux.NewRouter()

    r.HandleFunc("/student", student_handler.Register).Methods(http.MethodPost)
    r.HandleFunc("/student/{id}", student_handler.GetById).Methods(http.MethodGet)

    // Health Check
    r.HandleFunc("/health", public_handler.HealthCheck).Methods(http.MethodGet)

    port := ":8080"
    log.Printf("🚀 HTTP server running on http://localhost%s\n", port)

    if err := http.ListenAndServe(port, r); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}