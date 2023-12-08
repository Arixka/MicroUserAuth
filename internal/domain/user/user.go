package domain

//implementar con gorm
type User struct {
	ID       int
	Username string
	Email    string
	Password string //TODO! Considera almacenar una versión hash de la contraseña
}