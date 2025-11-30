package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpServer "user/micro/internal/adapters/http/server"
	mongoAdapter "user/micro/internal/adapters/persistence/mongo"
	"user/micro/internal/application/usecases"
	"user/micro/internal/domain/ports"
	"user/micro/internal/infra/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No se encontró .env, usando variables del sistema")
	}

	// Initialize application
	app, err := bootstrap()
	if err != nil {
		log.Fatalf("❌ Error durante el bootstrap: %v", err)
	}

	// Setup graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("Shutdown signal received")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := app.db.Client().Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
		os.Exit(0)
	}()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s", port)
	if err := app.router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// App holds all application components
type App struct {
	router *gin.Engine
	db     *mongo.Database
}

// bootstrap initializes all application components
func bootstrap() (*App, error) {
	// Initialize MongoDB connection (infrastructure layer)
	db, err, isProduction := config.ConnectDB()
	if err != nil {
		return nil, err
	}

	log.Println("Module connected to MongoDB")

	// Instantiate repository adapter (persistence layer)
	userRepository := mongoAdapter.NewUserRepository(db)

	// Verify database connection
	if err := userRepository.PingDataBase(); err != nil {
		return nil, err
	}

	log.Println("BBDD checked correctly")

	// Instantiate use cases (application layer)
	var userService ports.UserService = usecases.NewUserService(userRepository)

	// Setup HTTP adapter and routes
	router := httpServer.SetupRouter(userService, isProduction)

	return &App{
		router: router,
		db:     db,
	}, nil
}
