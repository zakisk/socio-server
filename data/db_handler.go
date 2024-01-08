package data

import (
	"fmt"

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

func (d *DBHandler) GetUserByCondition(key, value string) (*models.User, error) {
	var user models.User
	err := d.db.Where(fmt.Sprintf("%s = ?", key), &value).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *DBHandler) UpdateUser(user *models.User) error {
	return d.db.Save(user).Error
}

func (d *DBHandler) CreatePost(post *models.Post) error {
	return d.db.Create(post).Error
}

func (d *DBHandler) GetPosts() ([]models.Post, error) {
	posts := []models.Post{}
	err := d.db.Find(&posts).Error
	return posts, err
}
