package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakisk/socio-server/models"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.log.Info().Interface("User", user).Msg("got this values")
}

func (h *Handler) LoginUser(c *gin.Context) {

}
