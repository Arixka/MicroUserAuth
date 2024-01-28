package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/microservices/microUserAuth/internal/domain/user"
	"github.com/microservices/microUserAuth/internal/usecase/service"
	"golang.org/x/crypto/bcrypt"
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
	log.Printf("Entramos en Register  newUser '%+v':", newUser)
	log.Printf("Entramos en Register  newUser.Password '%+v':", newUser.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
		return
	}
	newUser.Password = string(hashedPassword)
	createdUser, err := h.userService.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Omitir la contraseña en la respuesta
	createdUser.Password = ""
	c.JSON(http.StatusCreated, createdUser)
}

// Aquí puedes agregar más métodos para actualizar, obtener y eliminar usuarios.
