package route

import (
	"user/micro/controllers"
	"user/micro/repository"
	"user/micro/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database, isProduction bool) *gin.Engine {

	if isProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userRoutes := router.Group("/users")
	{
		// userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.GetUsers)
		// userRoutes.GET("/:id", userController.GetUserByID)
		// userRoutes.PUT("/:id", userController.UpdateUser)
		// userRoutes.DELETE("/:id", userController.DeleteUser)
		userRoutes.GET("/ping", userController.PingDataBase)
		userRoutes.GET("/total", userController.GetTotalUsers)
	}

	return router
}
