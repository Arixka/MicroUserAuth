package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/microservices/microUserAuth/internal/domain/user"
	"github.com/microservices/microUserAuth/internal/usecase/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Método para iniciar sesión
func (h *AuthHandler) Login(c *gin.Context) {
	var credentials domain.Credentials
	// Respuesta en caso de error
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Llamada al servicio de autenticación
	user, err := h.authService.Login(credentials.Username, credentials.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	// Enviar respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

// Aquí puedes añadir otros métodos como Logout, ChangePassword, etc.

//********************************* EXPLICACION ***************************************

// Estructura AuthHandler: Esta estructura mantiene una referencia al servicio de autenticación (authService) que se utiliza para ejecutar la lógica de autenticación.
// Constructor NewAuthHandler: Crea una nueva instancia de AuthHandler. Aquí inyectamos la dependencia del servicio de autenticación.

// Método Login:
// - Vinculación de Datos (Binding):
//  Intentamos vincular los datos JSON de la solicitud HTTP a una estructura de Credentials (que debe ser definida en domain/user y contener Username y Password).
// - Uso del AuthService:
//	Llamamos al método Login del authService, pasando las credenciales. Este método se encargará de verificar si el usuario existe y si la contraseña es correcta.
// - Manejo de Respuestas:
//	Dependiendo del resultado de la autenticación, devolvemos una respuesta HTTP adecuada al cliente.

// Controladores para manejar las solicitudes HTTP. Esto incluye analizar los datos de la solicitud, validarlos y enviar respuestas HTTP.
// Punto de Entrada a la Lógica de la Aplicación: Actúa como el punto de entrada para las operaciones relacionadas con usuarios en la API. Convierte los datos de las solicitudes en un formato que la lógica de negocio puede entender y utiliza.
// Transformación de Datos: Puede involucrar transformar datos de la solicitud en estructuras o formatos requeridos por la lógica de negocio.
// Manejo de Respuestas: Formatea y devuelve las respuestas adecuadas y códigos de estado HTTP al cliente.
