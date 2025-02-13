package services

import (
	"99-user-service/internal/app/models"
	"99-user-service/internal/app/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers(offset int, limit int) ([]models.User, error) {
	return s.Repo.GetAllUsers(offset, limit)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}
