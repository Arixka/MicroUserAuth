package domain

import (
	"gorm.io/gorm"
)

//gorm.Model incluye campos estándar ID, CreatedAt, UpdatedAt, DeletedAt
type User struct {
	gorm.Model 
	Username string `gorm:"not null;unique_index:username" json:"username" form:"username"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null" json:"-"`
	//TODO! Considera almacenar una versión hash de la contraseña
}