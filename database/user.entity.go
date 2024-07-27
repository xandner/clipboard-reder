package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Phone    string `gorm:"unique" json:"phone"`
	Token    string `gorm:"unique" json:"token"`
}
