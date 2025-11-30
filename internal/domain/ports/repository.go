package ports

import (
	"user/micro/internal/domain/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=internal/domain/ports/repository.go -destination=mocks/mock_user_repository.go -package=mocks

// UserRepository defines the interface for user repository operations.
// This is the port that adapters must implement.
type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id primitive.ObjectID) (*models.User, error)
	Update(id primitive.ObjectID, user *models.User) error
	Delete(id primitive.ObjectID) error
	PingDataBase() error
	GetTotalUsers() (int64, error)
}
