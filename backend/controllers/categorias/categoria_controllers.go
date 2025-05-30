package controllers

import (
	"backend/clients"
	"backend/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategorias(ctx *gin.Context) {
	var categorias []models.Categoria
	if err := clients.DB.Find(&categorias).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las categorías"})
		return
	}
	ctx.JSON(http.StatusOK, categorias)
}

func GetCategoriaByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var categoria models.Categoria
	if err := clients.DB.First(&categoria, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Categoría no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, categoria)
}
