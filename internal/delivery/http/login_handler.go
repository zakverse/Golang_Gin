package http

import (
	"Golang_Gin/internal/domain"
	"Golang_Gin/internal/usecase"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginHandler struct {
	uc *usecase.LoginUseCase
}

func NewLoginHandler(r *gin.Engine, uc *usecase.LoginUseCase) {
	handler := &LoginHandler{uc}
	authRoot := r.Group("/auth")
	{
		authRoot.POST("/login", handler.Login)
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var login domain.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.uc.Login(login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "JWT_SECRET not configured",
		})
		return
	}

	claims := jwt.MapClaims{
		"user_id":  response.UserID,
		"username": response.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successfully",
		"user": gin.H{
			"id":       response.UserID,
			"username": response.Username,
		},
		"token": signedToken,
	})
}
