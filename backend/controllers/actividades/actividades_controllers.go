// Archivo: controllers/actividades_controller.go
package controllers

import (
	"backend/clients"
	"backend/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetActividadByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var actividad models.Actividad
	tx := clients.DB.Preload("Instructor").Preload("Categoria").First(&actividad, id)
	if tx.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Actividad no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, actividad)
}
