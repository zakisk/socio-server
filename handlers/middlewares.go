package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) AuthVerification(c *gin.Context) {
	c.Next()
}