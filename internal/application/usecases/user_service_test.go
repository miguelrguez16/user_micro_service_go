package usecases

import (
	"errors"
	"testing"

	"user/micro/internal/domain/models"
	"user/micro/mocks"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

// Test_UserService_FindAll_Success tests successful retrieval of all users
func Test_UserService_FindAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock using gomock
	mockRepo := mocks.NewMockUserRepository(ctrl)

	// Expected users
	expectedUsers := []models.User{
		{
			ID:    primitive.NewObjectID(),
			Name:  "John Doe",
			Email: "john@example.com",
		},
		{
			ID:    primitive.NewObjectID(),
			Name:  "Jane Smith",
			Email: "jane@example.com",
		},
	}

	// Setup expectations
	mockRepo.EXPECT().
		FindAll().
		Return(expectedUsers, nil).
		Times(1)

	// Create service with mock
	service := NewUserService(mockRepo)

	// Execute
	users, err := service.FindAll()

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(users) != len(expectedUsers) {
		t.Errorf("Expected %d users, got %d", len(expectedUsers), len(users))
	}
}

// Test_UserService_GetTotalUsers_Success tests successful retrieval of total users
func Test_UserService_GetTotalUsers_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	expectedTotal := int64(42)

	mockRepo.EXPECT().
		GetTotalUsers().
		Return(expectedTotal, nil).
		Times(1)

	service := NewUserService(mockRepo)

	total, err := service.GetTotalUsers()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if total != expectedTotal {
		t.Errorf("Expected %d users, got %d", expectedTotal, total)
	}
}

// Test_UserService_PingDataBase_Success tests successful database ping
func Test_UserService_PingDataBase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().
		PingDataBase().
		Return(nil).
		Times(1)

	service := NewUserService(mockRepo)

	result := service.PingDataBase()

	if !result {
		t.Errorf("Expected PingDataBase to return true, got false")
	}
}

// Test_UserService_PingDataBase_Failure tests failed database ping
func Test_UserService_PingDataBase_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().
		PingDataBase().
		Return(errors.New("database unavailable")).
		Times(1)

	service := NewUserService(mockRepo)

	result := service.PingDataBase()

	if result {
		t.Errorf("Expected PingDataBase to return false, got true")
	}
}

// Test_UserService_Create_Success tests successful user creation
func Test_UserService_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	user := &models.User{
		ID:    primitive.NewObjectID(),
		Name:  "John Doe",
		Email: "john@example.com",
	}

	mockRepo.EXPECT().
		Create(user).
		Return(nil).
		Times(1)

	service := NewUserService(mockRepo)

	err := service.Create(user)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test_UserService_FindByID_Success tests successful user retrieval by ID
func Test_UserService_FindByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	userID := primitive.NewObjectID()
	expectedUser := &models.User{
		ID:    userID,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	mockRepo.EXPECT().
		FindByID(userID).
		Return(expectedUser, nil).
		Times(1)

	service := NewUserService(mockRepo)

	user, err := service.FindByID(userID)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.ID != userID {
		t.Errorf("Expected user ID %v, got %v", userID, user.ID)
	}
}

// Test_UserService_Delete_Success tests successful user deletion
func Test_UserService_Delete_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	userID := primitive.NewObjectID()

	mockRepo.EXPECT().
		Delete(userID).
		Return(nil).
		Times(1)

	service := NewUserService(mockRepo)

	err := service.Delete(userID)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

// Test_UserService_Update_Success tests successful user update
func Test_UserService_Update_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	userID := primitive.NewObjectID()
	user := &models.User{
		ID:    userID,
		Name:  "John Updated",
		Email: "john.updated@example.com",
	}

	mockRepo.EXPECT().
		Update(userID, user).
		Return(nil).
		Times(1)

	service := NewUserService(mockRepo)

	err := service.Update(userID, user)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
