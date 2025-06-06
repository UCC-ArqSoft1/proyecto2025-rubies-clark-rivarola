// Archivo: controllers/actividades_controller.go
package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetActividadByID maneja GET /activities/:id y devuelve la actividad con todos sus horarios.
func GetActividadByID(ctx *gin.Context) {
	// 1. Parsear ID de la URL
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// 2. Llamar al servicio para obtener la actividad con sus horarios
	actividad, err := services.GetActivityByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// 3. Devolver JSON al cliente
	ctx.JSON(http.StatusOK, actividad)
}

// GetAllActivities maneja GET /activities y devuelve todas las actividades,
// opcionalmente filtradas por el query "search".
func GetAllActivities(ctx *gin.Context) {
	// Leer parámetro de búsqueda: ?search=<texto>
	query := ctx.Query("search")

	actividades, err := services.ListarActividades(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, actividades)
}
