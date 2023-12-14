package service

import (
	domain "github.com/microservices/microUserAuth/internal/domain/user"
)

// para crear el servicio necesitamos el modelo y el repository que interactua con la base de datos

//UserService: define las operaciones que se pueden realizar, pero no cómo se realizan
type UserService interface {
    CreateUser(user domain.User) (*domain.User, error)
    // Aquí puedes añadir más métodos como GetUserByID, UpdateUser, etc.
}

//es la definición de un tipo estructura en Go que implementará la interfaz UserService
type userServiceImpl struct {
    userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) UserService {
    return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) CreateUser(user domain.User) (*domain.User, error) {
    // Aquí puedes añadir lógica de negocio antes de guardar el usuario
    // Por ejemplo, validar datos del usuario, hash de contraseña, etc.

    createdUser, err := s.userRepo.CreateUser(user)
    if err != nil {
        return nil, err
    }

    return createdUser, nil
}

// Actualizar la Información del Usuario
// Consultar Usuarios
// Eliminar Usuarios