package http

import (
	"Golang_Gin/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BarangHandler struct {
	uc *usecase.BarangUseCase
}

func NewBarangHandler(r *gin.Engine, uc *usecase.BarangUseCase) {
	handler := &BarangHandler{uc}
	barangRoot := r.Group("/barang")
	{
		barangRoot.GET("/all", handler.GetAllBarang)
	}
}

func (h *BarangHandler) GetAllBarang(c *gin.Context) {
	barangs, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, barangs)
}
