package services_test

import (
	// "errors"
	"testing"

	"99-user-service/internal/app/models"
	"99-user-service/internal/app/repositories/mocks"
	"99-user-service/internal/app/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockUsers := []models.User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Doe"},
	}

	// Set expectation for GetAllUsers() in mock
	mockRepo.On("GetAllUsers", 0, 10).Return(mockUsers, nil)

	service := services.NewUserService(mockRepo)

	users, err := service.GetAllUsers(0, 10)
	assert.NoError(t, err)
	assert.Equal(t, mockUsers, users)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	newUser := &models.User{Name: "Alice"}

	mockRepo.On("CreateUser", mock.Anything).Return(nil)

	service := services.NewUserService(mockRepo)

	err := service.CreateUser(newUser)
	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "CreateUser", mock.Anything)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockUser := &models.User{ID: 1, Name: "John Doe"}

	// Set expectation for GetUserByID() in mock
	mockRepo.On("GetUserByID", uint(1)).Return(mockUser, nil)

	service := services.NewUserService(mockRepo)

	user, err := service.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)

	mockRepo.AssertExpectations(t)
}
