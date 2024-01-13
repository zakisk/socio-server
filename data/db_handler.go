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
	return d.db.Model(&models.User{}).
		Where("user_id = ?", user.UserID).
		Save(user).Error
}

func (d *DBHandler) CreatePost(post *models.Post) error {
	return d.db.Create(post).Error
}

func (d *DBHandler) GetPosts() ([]models.Post, error) {
	posts := []models.Post{}
	err := d.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, err
}

func (d *DBHandler) GetPostsByCondition(key, value string) ([]models.Post, error) {
	posts := []models.Post{}
	err := d.db.Where(fmt.Sprintf("%s = ?", key), value).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (d *DBHandler) GetPostByCondition(key, value string) (*models.Post, error) {
	var post models.Post
	err := d.db.Where(fmt.Sprintf("%s = ?", key), &value).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (d *DBHandler) UpdatePost(post models.Post) error {
	return d.db.Model(&models.Post{}).
		Where("post_id = ?", post.PostID).
		Save(post).Error
}
