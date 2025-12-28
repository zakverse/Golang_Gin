package http

import (
	"Golang_Gin/internal/domain"
	"Golang_Gin/internal/usecase"
	"net/http"
	"strconv"
	"Golang_Gin/internal/infrastructure"
	"github.com/gin-gonic/gin"
)

type BarangHandler struct {
	uc *usecase.BarangUseCase
}

func NewBarangHandler(r *gin.Engine, uc *usecase.BarangUseCase) {
	handler := &BarangHandler{uc}
	barangRoot := r.Group("/barang")
	{
		barangRoot.GET("/all", infrastructure.JWTAuth(), handler.GetAllBarang)
		barangRoot.GET("/:id", handler.GetByID)
		barangRoot.GET("/join/:id", handler.GetByIDJoin)
		barangRoot.GET("/join-detail-harga-sewa", handler.JoinDetailHargaSewa)
		barangRoot.POST("/create", handler.Create)
		barangRoot.PUT("/:id", handler.Update)
		barangRoot.DELETE("/:id", handler.Delete)
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

func (h *BarangHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	barang, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, barang)
}

func (h *BarangHandler) Create(c *gin.Context) {
	var barang domain.Barang
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(barang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Barang created successfully"})
}

func (h *BarangHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var barang domain.Barang
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorBagianJson": err.Error()})
		return
	}
	barang.ID = uint(id)
	if err := h.uc.Update(barang); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorBagianUpdate": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Barang updated successfully"})
}

func (h *BarangHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorBagianjson": "Invalid ID format"})
		return
	}
	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorBagianDelete": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Barang deleted successfully"})
}

func (h *BarangHandler) GetByIDJoin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	barang, err := h.uc.GetByIDJoin(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, barang)
}

func (h* BarangHandler) JoinDetailHargaSewa(c *gin.Context) {
	barang, err := h.uc.JoinDetailHargaSewa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, barang)
}