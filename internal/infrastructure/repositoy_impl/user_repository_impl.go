package repositoryimpl

import (
	"errors"
	"log"

	domain "github.com/microservices/microUserAuth/internal/domain/user"
	"github.com/microservices/microUserAuth/internal/usecase/service"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository crea una nueva instancia de userRepositoryImpl
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (repo *userRepositoryImpl) CreateUser(user domain.User) (*domain.User, error) {
	result := repo.db.Create(&user)
	log.Printf("Metodo CreateUser")
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindByUsername busca un usuario por su nombre de usuario
func (repo *userRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	u := &domain.User{}
	log.Printf("Entramos en FindByUsername '%s':", username)
	if err := repo.db.Where("username = ?", username).First(u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error al buscar el usuario '%s': %v", username, err)
			return nil, service.ErrInvalidCredentials
		}
		log.Printf("No se encontró el usuario '%s'", username)
		return nil, err
	}
	//Si un usuario es encontrado, se devuelve un puntero a la instancia del usuario (&u),
	log.Printf("Usuario encontrado: %v", u)
	return u, nil
}

func (repo *userRepositoryImpl) ResetPassword(userID uint, newPassword string) error {
	return nil
}
func (repo *userRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	u := &domain.User{}
	return u, nil
}
