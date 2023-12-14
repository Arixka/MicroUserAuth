package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/microservices/microUserAuth/internal/domain/user"
	"github.com/microservices/microUserAuth/internal/usecase/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var newUser domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.userService.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// Aquí puedes agregar más métodos para actualizar, obtener y eliminar usuarios.
