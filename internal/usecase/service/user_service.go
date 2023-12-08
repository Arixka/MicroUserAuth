package service

import domain "github.com/microservices/microUserAuth/internal/domain/user"

// para crear el servicio necesitamos el modelo y el repository que interactua con la base de datos

//UserService: define las operaciones que se pueden realizar, pero no cómo se realizan
type UserService interface {
    RegisterUser(user domain.User) error
    // GetUserByID, UpdateUser...
}
//es la definición de un tipo estructura en Go que implementará la interfaz UserService
type userService struct {
    userRepository UserRepository
}
func NewUser(/* dependencias */) UserService {
    return &userService{/* inicializa tus dependencias */}
}
// Crear un Nuevo Usuario
// Actualizar la Información del Usuario
// Consultar Usuarios
// Eliminar Usuarios

func (s *userService) RegisterUser(user User) error {
    // Lógica de negocio para registrar un usuario
    return nil
}