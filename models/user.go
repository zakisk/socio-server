package models

import "time"

type User struct {
	// ID of Post
	ID uint `json:"_id" gorm:"primaryKey;autoIncrement"`

	// first name of user
	FirstName string `json:"firstName" form:"firstName" binding:"required"`

	// last name of user
	LastName string `json:"lastName" form:"lastName" binding:"required"`

	// email of user
	Email string `json:"email" form:"email" binding:"required,email,unique"`

	// profile password
	Password string `json:"password" form:"password"`

	// path of user profile picture
	PicturePath string `json:"picturePath" form:"picturePath"`

	// friends of user
	Friends []string `json:"friends" gorm:"type:varchar(255)[]"`

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
