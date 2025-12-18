package domain

type Student struct {
    ID   int
    Name string
    Email string
}

// Port (interface) for persistence
type StudentRepository interface {
    Create(student Student) (Student, error)
    GetByID(id int) (Student, error)
}