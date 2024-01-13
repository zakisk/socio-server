package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) AuthVerification(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if !strings.Contains(token, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is not a valid token, doesn't contain `Bearer` keyword"})
		return
	}

	if err := verifyToken(token[7:]); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Next()
}

func verifyToken(tokenString string) error {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	if len(secretKey) == 0{
		return fmt.Errorf("JWT_SECRET is empty")
	}
	
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	return nil
}
