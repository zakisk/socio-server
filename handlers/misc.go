package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetImage(c *gin.Context) {
	c.File(fmt.Sprintf("./public/assets/%s", c.Param("imageName")))
}
