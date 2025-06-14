package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetActividadByID maneja GET /activities/:id y devuelve la actividad con todos sus horarios.
func GetActividadByID(ctx *gin.Context) {
	// Parsear ID de la URL
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llama al servicio para obtener la actividad con sus horarios
	actividad, err := services.GetActivityByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//Devulve JSON al cliente
	ctx.JSON(http.StatusOK, actividad)
}

// GetAllActivities maneja GET /activities y devuelve todas las actividades,
// opcionalmente filtradas por el query "search".
func GetAllActivities(ctx *gin.Context) {
	query := ctx.Query("search")
	list, err := services.ListarActividades(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}
