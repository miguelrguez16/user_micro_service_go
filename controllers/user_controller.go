package controllers

import (
	"net/http"
	"user/micro/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (ctrl *UserController) PingDataBase(c *gin.Context) {
	if ctrl.Service.PingDataBase() {
		c.JSON(http.StatusOK, gin.H{"message": "Conexión a la base de datos exitosa"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectar a la base de datos"})
	}
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.Service.UserRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}
