package handlers

import (
	"net/http"
	"strconv"

	models "Golang_Gin/Models"

	"github.com/gin-gonic/gin"
)

func GetBarangByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Barang harus berupa angka"})
		return
	}

	for _, barang := range models.DataBarang {
		if barang.ID == id {
			c.JSON(http.StatusOK, barang)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Barang tidak ditemukan"})
}
