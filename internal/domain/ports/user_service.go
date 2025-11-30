package ports

import (
	"user/micro/internal/domain/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=user_service.go -destination=mocks/mock_user_service.go -package=mocks

// UserService defines the interface for user service operations.
// This is the port that adapters (controllers) must depend on.
type UserService interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id primitive.ObjectID) (*models.User, error)
	Update(id primitive.ObjectID, user *models.User) error
	Delete(id primitive.ObjectID) error
	PingDataBase() bool
	GetTotalUsers() (int64, error)
}
