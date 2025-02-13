package repositories

import (
	"99-user-service/internal/app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers(offset int, limit int) ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetAllUsers(offset int, limit int) ([]models.User, error) {
	var users []models.User
	err := r.DB.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
