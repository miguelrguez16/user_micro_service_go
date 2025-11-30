package usecases

import (
	"errors"
	"testing"

	"user/micro/internal/domain/models"
	"user/micro/mocks"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

// Test_UserService_PingDataBase tests database connectivity with various scenarios
func Test_UserService_PingDataBase(t *testing.T) {
	tests := []struct {
		name           string
		mockError      error
		expectedResult bool
	}{
		{
			name:           "successful ping",
			mockError:      nil,
			expectedResult: true,
		},
		{
			name:           "failed ping",
			mockError:      errors.New("database unavailable"),
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				PingDataBase().
				Return(tt.mockError).
				Times(1)

			service := NewUserService(mockRepo)
			result := service.PingDataBase()

			if result != tt.expectedResult {
				t.Errorf("Expected PingDataBase to return %v, got %v", tt.expectedResult, result)
			}
		})
	}
}

// Test_UserService_GetTotalUsers tests GetTotalUsers with various scenarios
func Test_UserService_GetTotalUsers(t *testing.T) {
	tests := []struct {
		name        string
		mockReturn  int64
		mockError   error
		expectError bool
	}{
		{
			name:        "successful total with users",
			mockReturn:  42,
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "successful total with no users",
			mockReturn:  0,
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "database error",
			mockReturn:  0,
			mockError:   errors.New("connection failed"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				GetTotalUsers().
				Return(tt.mockReturn, tt.mockError).
				Times(1)

			service := NewUserService(mockRepo)
			total, err := service.GetTotalUsers()

			if (err != nil) != tt.expectError {
				t.Errorf("Expected error %v, got %v", tt.expectError, err != nil)
			}

			if !tt.expectError && total != tt.mockReturn {
				t.Errorf("Expected %d users, got %d", tt.mockReturn, total)
			}
		})
	}
}

// Test_UserService_FindAll tests FindAll with various scenarios
func Test_UserService_FindAll(t *testing.T) {
	tests := []struct {
		name          string
		mockUsers     []models.User
		mockError     error
		expectError   bool
		expectedCount int
	}{
		{
			name: "successful retrieval with multiple users",
			mockUsers: []models.User{
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
			},
			mockError:     nil,
			expectError:   false,
			expectedCount: 2,
		},
		{
			name:          "successful retrieval with no users",
			mockUsers:     []models.User{},
			mockError:     nil,
			expectError:   false,
			expectedCount: 0,
		},
		{
			name:          "database error",
			mockUsers:     nil,
			mockError:     errors.New("connection failed"),
			expectError:   true,
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				FindAll().
				Return(tt.mockUsers, tt.mockError).
				Times(1)

			service := NewUserService(mockRepo)
			users, err := service.FindAll()

			if (err != nil) != tt.expectError {
				t.Errorf("Expected error %v, got %v", tt.expectError, err != nil)
			}

			if len(users) != tt.expectedCount {
				t.Errorf("Expected %d users, got %d", tt.expectedCount, len(users))
			}
		})
	}
}

// Test_UserService_FindByID tests FindByID with various scenarios
func Test_UserService_FindByID(t *testing.T) {
	userID := primitive.NewObjectID()
	tests := []struct {
		name        string
		id          primitive.ObjectID
		mockUser    *models.User
		mockError   error
		expectError bool
	}{
		{
			name: "successful retrieval",
			id:   userID,
			mockUser: &models.User{
				ID:    userID,
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "user not found",
			id:          primitive.NewObjectID(),
			mockUser:    nil,
			mockError:   errors.New("user not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				FindByID(tt.id).
				Return(tt.mockUser, tt.mockError).
				Times(1)

			service := NewUserService(mockRepo)
			user, err := service.FindByID(tt.id)

			if (err != nil) != tt.expectError {
				t.Errorf("Expected error %v, got %v", tt.expectError, err != nil)
			}

			if !tt.expectError && user.ID != tt.id {
				t.Errorf("Expected user ID %v, got %v", tt.id, user.ID)
			}
		})
	}
}

// Test_UserService_Create tests Create with various scenarios
func Test_UserService_Create(t *testing.T) {
	tests := []struct {
		name        string
		user        *models.User
		mockError   error
		expectError bool
	}{
		{
			name: "successful creation",
			user: &models.User{
				ID:    primitive.NewObjectID(),
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockError:   nil,
			expectError: false,
		},
		{
			name: "creation with duplicate email",
			user: &models.User{
				ID:    primitive.NewObjectID(),
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockError:   errors.New("duplicate key error"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				Create(tt.user).
				Return(tt.mockError).
				Times(1)

			service := NewUserService(mockRepo)
			err := service.Create(tt.user)

			if (err != nil) != tt.expectError {
				t.Errorf("Expected error %v, got %v", tt.expectError, err != nil)
			}
		})
	}
}

// Test_UserService_Update tests Update with various scenarios
func Test_UserService_Update(t *testing.T) {
	userID := primitive.NewObjectID()
	tests := []struct {
		name        string
		id          primitive.ObjectID
		user        *models.User
		mockError   error
		expectError bool
	}{
		{
			name: "successful update",
			id:   userID,
			user: &models.User{
				Name:  "John Updated",
				Email: "john.updated@example.com",
			},
			mockError:   nil,
			expectError: false,
		},
		{
			name: "update non-existent user",
			id:   primitive.NewObjectID(),
			user: &models.User{
				Name:  "Jane Doe",
				Email: "jane@example.com",
			},
			mockError:   errors.New("user not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				Update(tt.id, tt.user).
				Return(tt.mockError).
				Times(1)

			service := NewUserService(mockRepo)
			err := service.Update(tt.id, tt.user)

			if (err != nil) != tt.expectError {
				t.Errorf("Expected error %v, got %v", tt.expectError, err != nil)
			}
		})
	}
}

// Test_UserService_Delete tests Delete with various scenarios
func Test_UserService_Delete(t *testing.T) {
	tests := []struct {
		name        string
		id          primitive.ObjectID
		mockError   error
		expectError bool
	}{
		{
			name:        "successful deletion",
			id:          primitive.NewObjectID(),
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "delete non-existent user",
			id:          primitive.NewObjectID(),
			mockError:   errors.New("user not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepository(ctrl)
			mockRepo.EXPECT().
				Delete(tt.id).
				Return(tt.mockError).
				Times(1)

			service := NewUserService(mockRepo)
			err := service.Delete(tt.id)

			if (err != nil) != tt.expectError {
				t.Errorf("Expected error %v, got %v", tt.expectError, err != nil)
			}
		})
	}
}
