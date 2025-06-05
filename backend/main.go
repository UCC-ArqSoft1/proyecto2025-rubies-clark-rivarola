// Archivo: main.go
package main

import (
	"backend/controllers"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializamos el router de Gin
	router := gin.New()

	// Middleware global: CORS y recuperación de pánico
	router.Use(utils.CORS, gin.Recovery())

	// -------------------------
	// Rutas de autenticación
	// -------------------------
	router.POST("/users/login", controllers.Login)

	// -------------------------
	// Rutas de actividades
	// -------------------------
	router.GET("/activities", controllers.GetAllActivities)
	router.GET("/activities/:id", controllers.GetActividadByID)
	router.GET("/activities/:id/horarios", controllers.GetHorariosPorActividad)

	// -------------------------
	// Rutas de horarios
	// -------------------------
	router.GET("/horarios/:id", controllers.GetHorarioByID)

	// -------------------------
	// Rutas de inscripciones
	// -------------------------
	router.POST("/inscripciones", controllers.CrearInscripcion)

	// -------------------------
	// Rutas de inscripciones por usuario
	// -------------------------
	router.GET("/usuarios/:id/inscripciones", controllers.GetInscripcionesPorUsuario)

	// Inicia el servidor en el puerto por defecto (":8080")
	router.Run()
}
