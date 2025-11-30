package http

import (
	"net/http"
	"user/micro/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

// HealthController handles health and ping requests for the microservice.
type HealthController struct {
	service ports.UserService
}

// NewHealthController creates a new HealthController instance.
func NewHealthController(service ports.UserService) *HealthController {
	return &HealthController{service: service}
}

// Ping returns the service name.
func (ctrl *HealthController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "user-microservice",
		"status":  "pong",
	})
}

// Health returns the health status of the microservice.
func (ctrl *HealthController) Health(c *gin.Context) {
	dbHealthy := ctrl.service.PingDataBase()

	status := "ok"
	dbStatus := "connected"
	statusCode := http.StatusOK

	if !dbHealthy {
		status = "degraded"
		dbStatus = "disconnected"
		statusCode = http.StatusServiceUnavailable
	}

	c.JSON(statusCode, gin.H{
		"status":   status,
		"service":  "user-microservice",
		"database": dbStatus,
	})
}
