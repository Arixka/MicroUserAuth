package repositoryimpl

import (
	domain "github.com/microservices/microUserAuth/internal/domain/user"
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
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// FindByUsername busca un usuario por su nombre de usuario
func (repo *userRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var u domain.User
	if err := repo.db.Where("username = ?", username).First(&u).Error; err != nil {
		// Manejar error, por ejemplo, si no se encuentra el usuario
		return nil, err
	}
	//Si un usuario es encontrado, se devuelve un puntero a la instancia del usuario (&u),
	return &u, nil
}
