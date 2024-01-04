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

	if err := h.dbHandler.InsertUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func (h *Handler) LoginUser(c *gin.Context) {
	
}
