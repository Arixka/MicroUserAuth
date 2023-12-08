package repositoryimpl

import (
	domain "github.com/microservices/microUserAuth/internal/domain/user"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}
//
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