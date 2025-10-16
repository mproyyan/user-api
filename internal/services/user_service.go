package services

import (
	"github.com/google/uuid"
	"github.com/mproyyan/user-api/internal/entities"
)

type UserService struct {
	repo entities.UserRepository
}

func NewUserService(repo entities.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id string) (*entities.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(name, email string) (*entities.User, error) {
	user := entities.User{
		ID:    uuid.NewString(),
		Name:  name,
		Email: email,
	}
	return s.repo.Create(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteByID(id)
}

func (s *UserService) IsEmailExists(email string) (bool, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}
