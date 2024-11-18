package domain

//definimos los metodos que interactuan con la base de datos, pero no su implementacion

type UserRepository interface {
	CreateUser(user User) (*User, error)
	FindByUsername(username string) (*User, error)
	ResetPassword(userID uint, newPassword string) error
	FindByEmail(email string) (*User, error)
	//FindUserByID, UpdateUser, DeleteUser, etc.
}
