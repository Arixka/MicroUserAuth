package domain

import (
	"gorm.io/gorm"
)

// gorm.Model incluye campos estándar ID, CreatedAt, UpdatedAt, DeletedAt
type User struct {
	gorm.Model
	Username string `gorm:"not null;unique_index:username" json:"username" form:"username"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null" json:"password"`
	//TODO! Considera almacenar una versión hash de la contraseña
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
