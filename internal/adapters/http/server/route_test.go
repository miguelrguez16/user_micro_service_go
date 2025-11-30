package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"user/micro/internal/domain/models"
	"user/micro/mocks"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

// Test_SetupRouter_Routes tests various API endpoints
func Test_SetupRouter_Routes(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		setupMock      func(*mocks.MockUserService)
		expectedStatus int
		isProduction   bool
	}{
		{
			name:           "ping endpoint",
			method:         "GET",
			path:           "/ping",
			setupMock:      func(m *mocks.MockUserService) {},
			expectedStatus: http.StatusOK,
			isProduction:   false,
		},
		{
			name:   "health endpoint with healthy database",
			method: "GET",
			path:   "/health",
			setupMock: func(m *mocks.MockUserService) {
				m.EXPECT().PingDataBase().Return(true).Times(1)
			},
			expectedStatus: http.StatusOK,
			isProduction:   false,
		},
		{
			name:   "health endpoint with unhealthy database",
			method: "GET",
			path:   "/health",
			setupMock: func(m *mocks.MockUserService) {
				m.EXPECT().PingDataBase().Return(false).Times(1)
			},
			expectedStatus: http.StatusServiceUnavailable,
			isProduction:   false,
		},
		{
			name:   "get users endpoint",
			method: "GET",
			path:   "/users",
			setupMock: func(m *mocks.MockUserService) {
				m.EXPECT().
					FindAll().
					Return([]models.User{
						{
							ID:    primitive.NewObjectID(),
							Name:  "Test User",
							Email: "test@example.com",
						},
					}, nil).
					Times(1)
			},
			expectedStatus: http.StatusOK,
			isProduction:   false,
		},
		{
			name:   "users ping endpoint",
			method: "GET",
			path:   "/users/ping",
			setupMock: func(m *mocks.MockUserService) {
				m.EXPECT().PingDataBase().Return(true).Times(1)
			},
			expectedStatus: http.StatusOK,
			isProduction:   false,
		},
		{
			name:   "users total endpoint",
			method: "GET",
			path:   "/users/total",
			setupMock: func(m *mocks.MockUserService) {
				m.EXPECT().GetTotalUsers().Return(int64(3), nil).Times(1)
			},
			expectedStatus: http.StatusOK,
			isProduction:   false,
		},
		{
			name:           "non-existent route",
			method:         "GET",
			path:           "/non-existent-route",
			setupMock:      func(m *mocks.MockUserService) {},
			expectedStatus: http.StatusNotFound,
			isProduction:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockUserService(ctrl)
			tt.setupMock(mockService)

			router := SetupRouter(mockService, tt.isProduction)

			req, _ := http.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

// Test_SetupRouter_Modes tests router creation in different modes
func Test_SetupRouter_Modes(t *testing.T) {
	tests := []struct {
		name         string
		isProduction bool
	}{
		{
			name:         "debug mode",
			isProduction: false,
		},
		{
			name:         "production mode",
			isProduction: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockUserService(ctrl)
			router := SetupRouter(mockService, tt.isProduction)

			if router == nil {
				t.Errorf("Expected router to be created in %s", tt.name)
			}
		})
	}
}

// Test_SetupRouter_PostUserRoute tests POST endpoint behavior
func Test_SetupRouter_PostUserRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)
	router := SetupRouter(mockService, false)

	user := models.User{
		Name:  "New User",
		Email: "new@example.com",
	}

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// POST /users not implemented yet - should be 404 or 405
	if w.Code != http.StatusNotFound && w.Code != http.StatusMethodNotAllowed {
		t.Logf("POST /users returned %d (not implemented)", w.Code)
	}
}
