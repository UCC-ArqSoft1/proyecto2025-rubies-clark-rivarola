// Archivo: controllers/inscripciones_controller.go
package controllers

import (
	"backend/clients"
	"backend/dao"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CrearInscripcion(ctx *gin.Context) {
	actividadIDStr := ctx.Param("id")
	actividadID, err := strconv.Atoi(actividadIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de actividad inválido"})
		return
	}

	var payload struct {
		UsuarioID uint `json:"usuario_id"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inscripcion := models.Inscripcion{
		UsuarioID:        payload.UsuarioID,
		ActividadID:      uint(actividadID),
		FechaInscripcion: time.Now(),
	}

	err = clients.DB.Create(&inscripcion).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la inscripción"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"mensaje": "Inscripción registrada con éxito"})
}
