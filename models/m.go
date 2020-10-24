package models

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	AccessToken  string
	RefreshToken string
}
