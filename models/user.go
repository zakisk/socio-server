package models

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {
	// UserID of Post
	UserID string `json:"_id"`

	// first name of user
	FirstName string `json:"firstName" form:"firstName" binding:"required"`

	// last name of user
	LastName string `json:"lastName" form:"lastName" binding:"required"`

	// email of user
	Email string `json:"email" form:"email" binding:"required,email" gorm:"unique"`

	// profile password
	Password string `json:"password" form:"password"`

	// path of user profile picture
	PicturePath string `json:"picturePath" form:"picturePath"`

	// friends of user
	Friends datatypes.JSON `json:"friends" gorm:"type:jsonb"`

	// location of user
	Location string `json:"location" form:"location"`

	// path user's profile picture
	Occupation string `json:"occupation" form:"occupation"`

	// likes on post
	ViewedProfile int `json:"viewedProfile" form:"viewedProfile"`

	// comments on post
	Impressions int `json:"impressions" form:"impressions"`

	// the time post is created
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`

	// the time post is updated
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}
