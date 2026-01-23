package domain

type Subject struct {
	ID int // Subject Code
	Name string
	Description string
}

// Port (interface) for persistence
type SubjectRepository interface {
	GetByID(id int) (Subject, error)
}