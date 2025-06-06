// Archivo: main.go
package main

import (
	"backend/controllers"
	database "backend/db"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	// Inicializar la base de datos (migraciones + seed)
	database.InitDB()

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

	// -------------------------
	// Rutas de inscripciones
	// -------------------------
	router.POST("/inscripciones", controllers.CrearInscripcion)

	// Inicia el servidor en el puerto por defecto (":8080")
	router.Run()
}
