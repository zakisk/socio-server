package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakisk/socio-server/helpers"
)

func (h *Handler) GetUser(c *gin.Context) {
	userId := c.Param("userId")
	if len(userId) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "userId in URL is empty"})
		return
	}

	user, err := h.dbHandler.GetUserByCondition(helpers.ID, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUserFriends(c *gin.Context) {
	
}

func (h *Handler) AddRemoveFriend(c *gin.Context) {

}
