package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Songs    []Song `gorm:"many2many:auth_songs;"`
}
