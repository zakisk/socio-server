package models

import "time"

type Post struct {
	// PostID of Post
	PostID string `json:"_id"`

	// ID of author user
	UserID string `json:"userId" form:"userId" binding:"required"`

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
	Likes map[string]bool `json:"likes" form:"likes" gorm:"type:json"`

	// comments on post
	Comments []string `json:"comments" form:"comments" gorm:"type:varchar(255)[]"`

	// the time post is created
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`

	// the time post is updated
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}

func (p *Post) Validate() {
	
}
