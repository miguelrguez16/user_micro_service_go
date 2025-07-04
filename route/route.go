package route

import (
	"user/micro/controllers"
	"user/micro/repository"
	"user/micro/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
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
	}

	return router
}
