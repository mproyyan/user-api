package entities

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id string) (*User, error)
	Create(user User) (*User, error)
	DeleteByID(id string) error
	FindByEmail(email string) (*User, error)
}
