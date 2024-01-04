package data

import (
	"github.com/zakisk/socio-server/models"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func NewDBHandler(db *gorm.DB) *DBHandler {
	db.AutoMigrate(models.User{}, models.Post{})
	return &DBHandler{DB: db}
}

func (d *DBHandler) InsertUser(user *models.User) error {
	return d.DB.Create(user).Error
}
