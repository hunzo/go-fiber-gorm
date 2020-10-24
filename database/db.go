package database

import (
	"github.com/hunzo/go-fiber-gorm/models"
	"gorm.io/gorm"
)

var DBcon *gorm.DB //db connect pointer

func Show() models.Token {
	// Initialize()
	db := DBcon
	var t models.Token
	db.Find(&t)

	// fmt.Println(t)
	return t
}
