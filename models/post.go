package models

import "time"

type Post struct {
	// ID of Post
	ID uint `json:"_id" gorm:"primaryKey;autoIncrement"`

	// ID of author user
	UserId string `json:"userId" form:"userId" binding:"required"`

	// first name of user
	FirstName string `json:"firstName" form:"firstName" binding:"required"`

	// last name of user
	LastName string `json:"lastName" form:"lastName" binding:"required"`

	// location of user
	Location string `json:"location" form:"location"`

	// description of post
	Description string `json:"description" form:"description"`

	// path of post image
	PicturePath string `json:"picturePath" form:"picturePath"`

	// path user's profile picture
	UserPicturePath string `json:"userPicturePath" form:"userPicturePath"`

	// likes on post
	Likes map[string]bool `json:"likes" form:"likes"`

	// comments on post
	Comments []string `json:"comments" form:"comments"`

	// the time post is created
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`

	// the time post is updated
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}

func (p *Post) Validate() {
	
}
