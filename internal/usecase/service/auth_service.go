package service

// Importar los paquetes necesarios
import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	domain "github.com/microservices/microUserAuth/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInternalServerError = errors.New("internal server error")

type AuthService interface {
	Login(username, password string) (string, error)
	// Aquí puedes añadir métodos como Logout, ChangePassword, etc.
}

type authServiceImpl struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
	}
}

func (s *authServiceImpl) Login(identifier, password string) (string, error) {
	var user *domain.User
	var err error
	if strings.Contains(identifier, "@") {
		user, err = s.userRepo.FindByUsername(identifier)
	} else {
		user, err = s.userRepo.FindByEmail(identifier)
	}
	if err != nil {
		log.Printf("Error al buscar el usuario '%s': %v", identifier, err)
		return "", err
	}
	if !s.checkPasswordHash(password, user.Password) {
		return "", ErrInvalidCredentials
	}
	token, err := s.generateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authServiceImpl) ResetPassword(email, password, newPassword string) (string, error) {
	return "null", nil
}

func (s *authServiceImpl) checkPasswordHash(password, hashedPassword string) bool {
	log.Printf("Metodo checkPasswordHash '%s': %v", password, hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("Error al verificar la contraseña: %v", err)
		return false
	}
	return true
}

type customClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (s *authServiceImpl) generateToken(user *domain.User) (string, error) {
	claims := &customClaims{
		Id:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Crear el token JWT firmado con HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Encapsula la logica de negocio

// authservice.go: Maneja la lógica de autenticación, incluyendo inicio y
// cierre de sesión, y posiblemente la gestión de tokens.

// Inicio de Sesión
// Cierre de Sesión
// Cambio de Contraseña
// Recuperación de Contraseña
