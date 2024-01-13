package handlers

import (
	"encoding/json"
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

	userFriends := []string{}
	err = json.Unmarshal(user.Friends, &userFriends)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, friendId := range userFriends {
		friend, _ := h.dbHandler.GetUserByCondition(helpers.UserID, friendId)
		friends = append(friends, *friend)
	}

	c.JSON(http.StatusOK, friends)
}

func (h *Handler) AddRemoveFriend(c *gin.Context) {
	id := c.Param("id")
	h.log.Info().Msg(fmt.Sprintf("user id is %s\n", id))
	if len(id) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "id in URL is empty"})
		return
	}

	friendId := c.Param("friendId")
	h.log.Info().Msg(fmt.Sprintf("friend id is %s\n", friendId))
	if len(friendId) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "id in URL is empty"})
		return
	}

	user, err := h.dbHandler.GetUserByCondition(helpers.UserID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("error while retrieving user: %s", err.Error())})
		return
	}

	friend, err := h.dbHandler.GetUserByCondition(helpers.UserID, friendId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("error while retrieving friend: %s", err.Error())})
		return
	}

	userFriends := []string{}
	if len(user.Friends) > 0 {
		err = json.Unmarshal(user.Friends, &userFriends)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	h.log.Info().Msg(fmt.Sprintf("user's friends before changes: %s\n", user.Friends))

	isFriend := false
	fIndex := 0
	for i, f := range userFriends {
		if f == friendId {
			isFriend = true
			fIndex = i
			break
		}
	}

	friendFriends := []string{}
	if len(friend.Friends) > 0 {
		err = json.Unmarshal(friend.Friends, &friendFriends)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	h.log.Info().Msg(fmt.Sprintf("friend's friends before changes: %s\n", friend.Friends))

	uIndex := 0
	for i, f := range friendFriends {
		if f == id {
			uIndex = i
			break
		}
	}

	if isFriend {
		userFriends = append(userFriends[:fIndex], userFriends[fIndex+1:]...)
		friendFriends = append(friendFriends[:uIndex], friendFriends[uIndex+1:]...)
	} else {
		userFriends = append(userFriends, friendId)
		friendFriends = append(friendFriends, id)
	}

	h.log.Info().Msg(fmt.Sprintf("user's friends after changes: %s\n", user.Friends))
	h.log.Info().Msg(fmt.Sprintf("friend's friends after changes: %s\n", friend.Friends))

	data, err := json.Marshal(userFriends)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Friends = data

	err = h.dbHandler.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error while updating user: %s", err.Error())})
		return
	}

	data, err = json.Marshal(friendFriends)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	friend.Friends = data

	err = h.dbHandler.UpdateUser(friend)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error while updating friend: %s", err.Error())})
		return
	}

	friends := []models.User{}
	for _, friendId := range userFriends {
		friend, _ := h.dbHandler.GetUserByCondition(helpers.UserID, friendId)
		friends = append(friends, *friend)
	}

	c.JSON(http.StatusOK, friends)
}
