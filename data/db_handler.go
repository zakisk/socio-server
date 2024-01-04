package data

import (
	"github.com/zakisk/socio-server/models"
	"gorm.io/gorm"
)

type DBHandler struct {
	db *gorm.DB
}

func NewDBHandler(db *gorm.DB) *DBHandler {
	db.AutoMigrate(models.User{}, models.Post{})
	return &DBHandler{db: db}
}

func (d *DBHandler) InsertUser(user *models.User) error {
	return d.db.Create(user).Error
}

func (d *DBHandler) GetUser() {

}
