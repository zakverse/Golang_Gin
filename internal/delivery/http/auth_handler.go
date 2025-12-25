package http

import (
	"Golang_Gin/internal/domain"
	"Golang_Gin/internal/usecase"
	"net/http"
	//"strconv"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc *usecase.AuthUseCase
}

func NewAuthHandler(r *gin.Engine, uc *usecase.AuthUseCase) {
	handler := &AuthHandler{uc}
	authRoot := r.Group("/auth")
	{
		authRoot.POST("/register",handler.Register)
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var register domain.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Register(register); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Aunth successfully"})
}

