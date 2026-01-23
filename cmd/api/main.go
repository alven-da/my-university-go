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
    subject_repo := repository.NewSubjectMemoryRepo()
    
    // initialize services
    student_service := usecase.NewStudentService(repo)
    public_service := usecase.NewPublicService()
    subject_service := usecase.NewSubjectService(subject_repo)

    // initialize http handlers
    student_handler := http_handler.NewStudentHandler(student_service)
    public_handler := http_handler.NewPublicHandler(public_service)
    subject_handler := http_handler.NewSubjectHandler(subject_service)

    r := mux.NewRouter()

    // Student API
    r.HandleFunc("/students", student_handler.Register).Methods(http.MethodPost)
    r.HandleFunc("/students/{id}", student_handler.GetById).Methods(http.MethodGet)

    // Health Check API
    r.HandleFunc("/health", public_handler.HealthCheck).Methods(http.MethodGet)

    // Subject API
    r.HandleFunc("/subjects/{id}", subject_handler.GetSubjectByID).Methods(http.MethodGet)

    port := ":8080"
    log.Printf("🚀 HTTP server running on http://localhost%s\n", port)

    if err := http.ListenAndServe(port, r); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}