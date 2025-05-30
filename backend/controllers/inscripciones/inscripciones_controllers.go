// Archivo: controllers/instructores_controller.go
package controllers

import (
	"backend/clients"
	"backend/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllInstructores(ctx *gin.Context) {
	var instructores []models.Instructor
	if err := clients.DB.Find(&instructores).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los instructores"})
		return
	}
	ctx.JSON(http.StatusOK, instructores)
}

func GetInstructorByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var instructor models.Instructor
	if err := clients.DB.First(&instructor, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Instructor no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, instructor)
}
