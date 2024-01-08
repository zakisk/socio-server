package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zakisk/socio-server/models"
)

type Router struct {
	S *gin.Engine
}

func NewRouter(handler models.HandlerInterface) *Router {
	engine := gin.New()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	corsMiddleware := cors.New(config)
	engine.Use(gin.Logger(), corsMiddleware)

	engine.POST("/auth/register", handler.RegisterUser)
	engine.POST("/auth/login", handler.LoginUser)

	usersGroup := engine.Group("/users")
	usersGroup.Use(handler.AuthVerification)
	{
		// Get
		usersGroup.GET("/:userId", handler.GetUser)
		usersGroup.GET("/:userId/friends", handler.GetUserFriends)

		// Patch
		usersGroup.PATCH("/:id/:friendId", handler.AddRemoveFriend)
	}

	postsGroup := engine.Group("/posts", handler.AuthVerification)
	{
		postsGroup.POST("/", handler.CreatePost)
		postsGroup.GET("/", handler.GetFeedPosts)
		postsGroup.GET("/:userId/posts", handler.GetUserPosts)
		postsGroup.PATCH("/:id/like", handler.LikePost)
	}

	engine.GET("/assets/:imageName", handler.GetImage)

	return &Router{S: engine}
}
