package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zakisk/socio-server/models"
)

type Router struct {
	S *gin.Engine
}

func NewRouter(handler models.HandlerInterface) *Router {
	engine := gin.New()

	engine.POST("/auth/register", handler.RegisterUser)
	engine.POST("/auth/login", handler.LoginUser)

	usersGroup := engine.Group("/users", handler.AuthVerification)
	{
		// Get
		usersGroup.GET("/:userId", handler.GetUser)
		usersGroup.GET("/:userId/friends", handler.GetUserFriends)

		// Patch
		usersGroup.PATCH("/:id/:friendId", handler.AddRemoveFriend)
	}

	postsGroup := engine.Group("/posts", handler.AuthVerification)
	{
		postsGroup.POST("/", handler.AuthVerification)
		postsGroup.GET("/", handler.GetFeedPosts)
		postsGroup.GET("/:userId/post", handler.GetUserPosts)
		postsGroup.PATCH("/:id/like", handler.LikePost)
	}

	return &Router{S: engine}
}
