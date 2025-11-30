package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"user/micro/internal/domain/models"
	"user/micro/mocks"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

// Test_HealthController tests HealthController with various scenarios
func Test_HealthController(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		dbHealthy      bool
		expectedStatus int
		hasBody        bool
	}{
		{
			name:           "ping endpoint returns ok",
			method:         "ping",
			dbHealthy:      true,
			expectedStatus: http.StatusOK,
			hasBody:        true,
		},
		{
			name:           "health endpoint with healthy database",
			method:         "health",
			dbHealthy:      true,
			expectedStatus: http.StatusOK,
			hasBody:        true,
		},
		{
			name:           "health endpoint with unhealthy database",
			method:         "health",
			dbHealthy:      false,
			expectedStatus: http.StatusServiceUnavailable,
			hasBody:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockUserService(ctrl)

			if tt.method == "health" {
				mockService.EXPECT().PingDataBase().Return(tt.dbHealthy).Times(1)
			}

			controller := NewHealthController(mockService)

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if tt.method == "ping" {
				controller.Ping(c)
			} else {
				controller.Health(c)
			}

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.hasBody && w.Body.String() == "" {
				t.Errorf("Expected response body, got empty")
			}
		})
	}
}

// Test_UserController_GetUsers tests GetUsers endpoint with various scenarios
func Test_UserController_GetUsers(t *testing.T) {
	tests := []struct {
		name           string
		mockUsers      []models.User
		mockError      error
		expectedStatus int
		hasBody        bool
	}{
		{
			name: "successful get users",
			mockUsers: []models.User{
				{
					ID:    primitive.NewObjectID(),
					Name:  "Test User",
					Email: "test@example.com",
				},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			hasBody:        true,
		},
		{
			name:           "database error",
			mockUsers:      nil,
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
			hasBody:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockUserService(ctrl)
			mockService.EXPECT().
				FindAll().
				Return(tt.mockUsers, tt.mockError).
				Times(1)

			controller := NewUserController(mockService)

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			controller.GetUsers(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.hasBody && w.Body.String() == "" {
				t.Errorf("Expected response body, got empty")
			}
		})
	}
}

// Test_UserController_GetTotalUsers tests GetTotalUsers endpoint with various scenarios
func Test_UserController_GetTotalUsers(t *testing.T) {
	tests := []struct {
		name           string
		mockTotal      int64
		mockError      error
		expectedStatus int
		hasBody        bool
	}{
		{
			name:           "successful get total users",
			mockTotal:      5,
			mockError:      nil,
			expectedStatus: http.StatusOK,
			hasBody:        true,
		},
		{
			name:           "get total with zero users",
			mockTotal:      0,
			mockError:      nil,
			expectedStatus: http.StatusOK,
			hasBody:        true,
		},
		{
			name:           "database error",
			mockTotal:      0,
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
			hasBody:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockUserService(ctrl)
			mockService.EXPECT().
				GetTotalUsers().
				Return(tt.mockTotal, tt.mockError).
				Times(1)

			controller := NewUserController(mockService)

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			controller.GetTotalUsers(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.hasBody && w.Body.String() == "" {
				t.Errorf("Expected response body, got empty")
			}
		})
	}
}

// Test_UserController_PingDataBase tests PingDataBase endpoint with various scenarios
func Test_UserController_PingDataBase(t *testing.T) {
	tests := []struct {
		name           string
		dbHealthy      bool
		expectedStatus int
	}{
		{
			name:           "successful database ping",
			dbHealthy:      true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "failed database ping",
			dbHealthy:      false,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockUserService(ctrl)
			mockService.EXPECT().
				PingDataBase().
				Return(tt.dbHealthy).
				Times(1)

			controller := NewUserController(mockService)

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			controller.PingDataBase(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
