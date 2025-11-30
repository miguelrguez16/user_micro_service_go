package usecases

import (
	"user/micro/internal/domain/models"
	"user/micro/internal/domain/ports"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserService implements the ports.UserService interface for user operations.
type UserService struct {
	userRepository ports.UserRepository
}

// NewUserService creates a new UserService instance.
func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &UserService{userRepository: userRepository}
}

// Create creates a new user.
func (us *UserService) Create(user *models.User) error {
	return us.userRepository.Create(user)
}

// FindAll retrieves all users.
func (us *UserService) FindAll() ([]models.User, error) {
	return us.userRepository.FindAll()
}

// FindByID retrieves a user by ID.
func (us *UserService) FindByID(id primitive.ObjectID) (*models.User, error) {
	return us.userRepository.FindByID(id)
}

// Update updates a user.
func (us *UserService) Update(id primitive.ObjectID, user *models.User) error {
	return us.userRepository.Update(id, user)
}

// Delete deletes a user.
func (us *UserService) Delete(id primitive.ObjectID) error {
	return us.userRepository.Delete(id)
}

// PingDataBase checks the connection to the database.
func (us *UserService) PingDataBase() bool {
	return us.userRepository.PingDataBase() == nil
}

// GetTotalUsers returns the total count of users.
func (us *UserService) GetTotalUsers() (int64, error) {
	return us.userRepository.GetTotalUsers()
}
