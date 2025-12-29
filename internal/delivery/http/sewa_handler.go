package http

import (
	"Golang_Gin/internal/domain"
	"Golang_Gin/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SewaHandler struct {
	uc *usecase.SewaUseCase
}

func NewSewaHandler(r *gin.Engine, uc *usecase.SewaUseCase) {
	handler := &SewaHandler{uc}
	authRoot := r.Group("/auth")
	{
		authRoot.POST("/sewa", handler.Create)
	}
}

func (h *SewaHandler) Create(c *gin.Context) {
	var sewa domain.SewaBody
	if err := c.ShouldBindJSON(&sewa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(sewa); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Sewa created successfully",
		"DetailSewaBody": 		 sewa.DetailSewaBody,
		"IDUser":           	 sewa.IDUser,
		"CodeStruk":        	 sewa.CodeStruk,
		"TanggalPeminjaman":	 sewa.TanggalPeminjaman,
		"TanggalJatuhTempo":	 sewa.TanggalJatuhTempo,
		"TanggalPengembalian":	 sewa.TanggalPengembalian,
		"PointReward":       	 sewa.PointReward,
		"DendaKeterlambatan":	 sewa.DendaKeterlambatan,
	})
}
