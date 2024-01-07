package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) AuthVerification(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if !strings.Contains(token, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is not a valid token, doesn't contain Bearer keyword"})
		return
	}

	if ok := verifyToken(token); !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	c.Next()
}

func verifyToken(tokenString string) bool {
	secretKey := []byte(os.Getenv("KWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
