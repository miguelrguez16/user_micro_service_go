package server

import (
	"user/micro/internal/adapters/http"
	"user/micro/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the Gin router with all routes.
func SetupRouter(userService ports.UserService, isProduction bool) *gin.Engine {
	if isProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	gin.Recovery()
	router := gin.Default()

	userController := http.NewUserController(userService)
	healthController := http.NewHealthController(userService)

	// Setup no route
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Path not found"})
	})

	// Health and ping endpoints
	router.GET("/ping", healthController.Ping)
	router.GET("/health", healthController.Health)

	// User endpoints
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", userController.GetUsers)
		userRoutes.GET("/ping", userController.PingDataBase)
		userRoutes.GET("/total", userController.GetTotalUsers)
	}

	return router
}
