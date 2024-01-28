package service

// Importar los paquetes necesarios
import (
	"errors"
	"log"

	domain "github.com/microservices/microUserAuth/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

// ErrInvalidCredentials se utiliza cuando un intento de login tiene credenciales inválidas
var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService interface {
	Login(username, password string) (*domain.User, error)
	// Aquí puedes añadir métodos como Logout, ChangePassword, etc.
}

type authServiceImpl struct {
	// Aquí irían las dependencias, como un repositorio de usuarios
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
	}
}

func (s *authServiceImpl) Login(username, password string) (*domain.User, error) {
	// Implementa la lógica de inicio de sesión aquí
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		log.Printf("Error al buscar el usuario '%s': %v", username, err)
		return nil, err
	}
	// Verifica la contraseña (asegúrate de que esté hasheada)
	if !s.checkPasswordHash(password, user.Password) {
		return nil, ErrInvalidCredentials
	}
	return user, nil
}

// checkPasswordHash compara una contraseña con un hash y devuelve si son iguales
func (s *authServiceImpl) checkPasswordHash(password, hashedPassword string) bool {
	log.Printf("Metodo checkPasswordHash '%s': %v", password, hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("Error al verificar la contraseña: %v", err)
		return false
	}
	return true
}

// Encapsula la logica de negocio

// authservice.go: Maneja la lógica de autenticación, incluyendo inicio y
// cierre de sesión, y posiblemente la gestión de tokens.

// Inicio de Sesión
// Cierre de Sesión
// Cambio de Contraseña
// Recuperación de Contraseña
