package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zakisk/socio-server/helpers"
	"github.com/zakisk/socio-server/models"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashedPassword
	user.UserID = helpers.GenerateUUID()

	if err := h.dbHandler.InsertUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func (h *Handler) LoginUser(c *gin.Context) {
	var body loginBody
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.dbHandler.GetUserByCondition(helpers.Email, body.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = helpers.CheckPassword(body.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("email or password is incorrect : %s", err.Error())})
		return
	}

	token, err := createToken(user.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &struct {
		User  *models.User `json:"user"`
		Token string       `json:"token"`
	}{User: user, Token: token})
}

type loginBody struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

func createToken(email string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (24 hours in this case)
	})

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	if len(secretKey) == 0 {
		return "", fmt.Errorf("JWT_SECRET is empty")
	}
	
	return claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
