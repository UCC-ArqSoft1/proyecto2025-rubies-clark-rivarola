package controllers

import (
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InscripcionRequest struct {
	UsuarioID uint `json:"usuario_id"`
	HorarioID uint `json:"horario_id"`
}

func CrearInscripcion(ctx *gin.Context) {
	var req InscripcionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if req.UsuarioID == 0 || req.HorarioID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "usuario_id y horario_id son obligatorios"})
		return
	}

	err := services.CrearInscripcion(req.UsuarioID, req.HorarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"mensaje": "Inscripción creada exitosamente"})
}

// POST /inscripciones/:horario_id
// Crea una inscripción del usuario autenticado a un horario específico.
func PostInscripcion(ctx *gin.Context) {
	// Obtiene ID del horario desde la URL
	horarioIDStr := ctx.Param("horario_id")
	horarioID, err := strconv.Atoi(horarioIDStr)
	if err != nil || horarioID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de horario inválido"})
		return
	}

	// Obtiene el ID del usuario autenticado
	userIDAny, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No autorizado"})
		return
	}
	userID := userIDAny.(uint)

	// Intentar crear la inscripción
	err = services.CrearInscripcion(userID, uint(horarioID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Inscripción realizada con éxito"})
}

// GET /usuarios/:usuario_id/inscripciones
// Lista todas las inscripciones del usuario (sus horarios de clase)
func GetInscripcionesPorUsuario(ctx *gin.Context) {
	idStr := ctx.Param("usuario_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	horarios, err := services.ListarInscripcionesPorUsuario(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, horarios)
}
