package models

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	//Middleware
	AuthVerification(c *gin.Context)
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)

	GetUser(c *gin.Context)
	GetUserFriends(c *gin.Context)
	AddRemoveFriend(c *gin.Context)

	CreatePost(c *gin.Context)
	GetFeedPosts(c *gin.Context)
	GetUserPosts(c *gin.Context)
	LikePost(c *gin.Context)

	GetImage(c *gin.Context)
}
