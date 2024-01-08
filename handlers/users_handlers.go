package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakisk/socio-server/helpers"
	"github.com/zakisk/socio-server/models"
)

func (h *Handler) GetUser(c *gin.Context) {
	userId := c.Param("userId")
	if len(userId) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "userId in URL is empty"})
		return
	}

	user, err := h.dbHandler.GetUserByCondition(helpers.UserID, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUserFriends(c *gin.Context) {
	userId := c.Param("userId")
	if len(userId) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "userId in URL is empty"})
		return
	}

	friends := []models.User{}
	user, err := h.dbHandler.GetUserByCondition(helpers.UserID, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	for _, friendId := range user.Friends {
		friend, _ := h.dbHandler.GetUserByCondition(helpers.UserID, friendId)
		friends = append(friends, *friend)
	}

	c.JSON(http.StatusOK, friends)
}

func (h *Handler) AddRemoveFriend(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "id in URL is empty"})
		return
	}

	friendId := c.Param("friendId")
	if len(friendId) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "id in URL is empty"})
		return
	}

	user, err := h.dbHandler.GetUserByCondition(helpers.UserID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("error while retrieving user: %s", err.Error())})
		return
	}

	friend, err := h.dbHandler.GetUserByCondition(helpers.UserID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("error while retrieving friend: %s", err.Error())})
		return
	}

	isFriend := false
	fIndex := 0
	for i, f := range user.Friends {
		if id == f {
			isFriend = true
			fIndex = i
			break
		}
	}

	uIndex := 0
	for i, f := range friend.Friends {
		if id == f {
			uIndex = i
			break
		}
	}

	if isFriend {
		user.Friends = append(user.Friends[:fIndex], user.Friends[fIndex+1:]...)
		friend.Friends = append(friend.Friends[:uIndex], friend.Friends[uIndex+1:]...)
	} else {
		user.Friends = append(user.Friends, friendId)
		friend.Friends = append(friend.Friends, id)
	}

	err = h.dbHandler.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error while updating user: %s", err.Error())})
		return
	}

	err = h.dbHandler.UpdateUser(friend)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error while updating user: %s", err.Error())})
		return
	}

	friends := []models.User{}
	for _, friendId := range user.Friends {
		friend, _ := h.dbHandler.GetUserByCondition(helpers.UserID, friendId)
		friends = append(friends, *friend)
	}

	c.JSON(http.StatusOK, friends)
}
