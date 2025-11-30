package http

import (
	"net/http"
	"user/micro/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

// UserController handles HTTP requests for user operations.
type UserController struct {
	service ports.UserService
}

// NewUserController creates a new UserController instance.
func NewUserController(service ports.UserService) *UserController {
	return &UserController{service: service}
}

// PingDataBase checks the database connection.
func (ctrl *UserController) PingDataBase(c *gin.Context) {
	if ctrl.service.PingDataBase() {
		c.JSON(http.StatusOK, gin.H{"message": "Conexión a la base de datos exitosa"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar a la base de datos"})
	}
}

// GetUsers retrieves all users.
func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetTotalUsers retrieves the total count of users.
func (ctrl *UserController) GetTotalUsers(c *gin.Context) {
	totalUsers, err := ctrl.service.GetTotalUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el total de usuarios"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_users": totalUsers})
}
