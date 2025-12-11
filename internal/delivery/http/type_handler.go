package http

import (
	"Golang_Gin/internal/domain"
	"Golang_Gin/internal/usecase"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TypeHandler struct {
	uc *usecase.TypeUseCase
}

func NewTypeHandler(r *gin.Engine, uc *usecase.TypeUseCase) {
	handler := &TypeHandler{uc}
	typeroot := r.Group("/types")
	{
		typeroot.GET("/all", handler.GetAllType)
		typeroot.GET("/:zaki/:abang", handler.GetByIDType)
		typeroot.POST("/create", handler.CreateType)
		typeroot.PUT("/:id", handler.UpdateType)
		typeroot.DELETE("/:id", handler.DeleteType)
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

func (h *TypeHandler) GetByIDType(c *gin.Context) {
	idStr := c.Param("zaki")
	abangstr := c.Param("abang")
	queryNgasal := c.Query("ngasalQuery")
	queryNgasal2 := c.Query("ngasalBanget")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	typ, err := h.uc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// c.JSON(http.StatusOK, typ)
	c.JSON(http.StatusOK, gin.H{
		"Dzaki" : "Mantap",
		"Abang" : abangstr,
		"Id_Atas" : typ.ID,
		"Data_all" : typ,
		"NgasalQuery" : queryNgasal,
		"NgasalBanget" : queryNgasal2,
	})
}

func (h *TypeHandler) CreateType(c *gin.Context) {
	var typ domain.Type
	if err := c.ShouldBindJSON(&typ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(typ); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Type created successfully"})
}

func (h *TypeHandler) UpdateType(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var typ domain.Type
	if err := c.ShouldBindJSON(&typ); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	typ.ID = uint(id)
	if err := h.uc.Update(typ); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Type updated successfully"})
}
func (h *TypeHandler) DeleteType(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.uc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Type deleted successfully"})
}
