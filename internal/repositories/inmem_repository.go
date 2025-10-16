package repositories

import (
	"errors"

	"github.com/mproyyan/user-api/internal/entities"
)

type InMemoryUserRepository struct {
	data map[string]entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: make(map[string]entities.User),
	}
}

func (r *InMemoryUserRepository) FindAll() ([]entities.User, error) {
	users := make([]entities.User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}

	return users, nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*entities.User, error) {
	user, exists := r.data[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *InMemoryUserRepository) Create(user entities.User) (*entities.User, error) {
	r.data[user.ID] = user
	return &user, nil
}

func (r *InMemoryUserRepository) DeleteByID(id string) error {
	if _, exists := r.data[id]; !exists {
		return errors.New("user not found")
	}
	
	delete(r.data, id)
	return nil
}
