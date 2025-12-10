package http

import (
	// "Golang_Gin/internal/domain"
	"Golang_Gin/internal/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

type TypeHandler struct{
	uc *usecase.TypeUseCase
}



func NewTypeHandler(r *gin.Engine, uc *usecase.TypeUseCase){
	handler := &TypeHandler{uc}
	typeroot :=  r.Group("/types") 
	{
		typeroot.GET("/all", handler.GetAllType)
	}
}

func (h *TypeHandler) GetAllType(c *gin.Context) {
	types, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, types)

}
