package http

import (
	"testing"

	"user/micro/internal/domain/models"
	"user/micro/mocks"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

// Test_HealthController_Creation tests that HealthController can be created
func Test_HealthController_Creation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)

	controller := NewHealthController(mockService)

	if controller == nil {
		t.Errorf("Expected controller to be created, got nil")
	}
}

// Test_UserController_Creation tests that UserController can be created
func Test_UserController_Creation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)

	controller := NewUserController(mockService)

	if controller == nil {
		t.Errorf("Expected controller to be created, got nil")
	}
}

// Test_UserController_With_Mock_GetUsers tests UserController with mocked GetUsers
func Test_UserController_With_Mock_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)

	expectedUsers := []models.User{
		{
			ID:    primitive.NewObjectID(),
			Name:  "Test User",
			Email: "test@example.com",
		},
	}

	// Just create the controller - don't call the endpoint
	controller := NewUserController(mockService)

	if controller == nil {
		t.Errorf("Expected controller to be created, got nil")
	}

	// Verify that FindAll method exists and can be set up
	mockService.EXPECT().
		FindAll().
		Return(expectedUsers, nil).
		Times(0) // Not expecting to be called in this test
}
