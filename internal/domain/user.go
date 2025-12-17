// internal/domain/user.go
package domain

type User struct {
    ID   int
    Name string
    Email string
}

// Port (interface) for persistence
type UserRepository interface {
    Create(user User) (User, error)
    GetByID(id int) (User, error)
}