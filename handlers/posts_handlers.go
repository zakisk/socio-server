package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakisk/socio-server/helpers"
	"github.com/zakisk/socio-server/models"
	"gorm.io/datatypes"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var body postBody
	if err := c.BindJSON(&body); err != nil {
		h.log.Info().Msg("Error while reading body")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.dbHandler.GetUserByCondition(helpers.UserID, body.UserId)
	if err != nil {
		h.log.Info().Msg("Error while retrieving user")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	post := &models.Post{
		PostID:          helpers.GenerateUUID(),
		UserID:          body.UserId,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Location:        user.Location,
		Description:     body.Description,
		PicturePath:     body.PicturePath,
		UserPicturePath: user.PicturePath,
		Likes:           datatypes.JSON(""),
		Comments:        []string{},
	}

	err = h.dbHandler.CreatePost(post)
	if err != nil {
		h.log.Info().Msg("Error creating post")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	posts, err := h.dbHandler.GetPosts()
	if err != nil {
		h.log.Info().Msg("Error while retrieving posts")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetFeedPosts(c *gin.Context) {
	posts, err := h.dbHandler.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetUserPosts(c *gin.Context) {
	userId := c.Param("userId")
	if len(userId) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "userId in URL is empty"})
		return
	}

	posts, err := h.dbHandler.GetPostsByCondition(helpers.UserID, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("error while retrieving user's posts: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) LikePost(c *gin.Context) {
	var body likeBody
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "id in URL is empty"})
		return
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.dbHandler.GetPostByCondition(helpers.PostID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	m := map[string]bool{}
	if len(post.Likes) > 0 {
		err = json.Unmarshal(post.Likes, &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if ok := m[body.UserId]; ok {
		delete(m, body.UserId)
	} else {
		m[body.UserId] = true
	}

	data, err := json.Marshal(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post.Likes = data
	err = h.dbHandler.UpdatePost(*post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

type postBody struct {
	UserId      string `json:"userId" form:"userId" binding:"required"`
	Description string `json:"description" form:"description"`
	PicturePath string `json:"picturePath" form:"picturePath" binding:"required"`
}

type likeBody struct {
	UserId string `json:"userId" form:"userId" binding:"required"`
}
