package main

import (
	"backend/controllers"
	"backend/services"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	services.Login("emiliano", "1234")

	router := gin.New()

	// Middleware CORS para cada ruta
	router.POST("/users/login", utils.CORS, controllers.Login)

	// Actividades
	router.GET("/activities/:id", utils.CORS, controllers.GetActividadByID)

	// Categorías
	router.GET("/categorias", utils.CORS, controllers.GetAllCategorias)
	router.GET("/categorias/:id", utils.CORS, controllers.GetCategoriaByID)

	// Instructores
	router.GET("/instructores", utils.CORS, controllers.GetAllInstructores)
	router.GET("/instructores/:id", utils.CORS, controllers.GetInstructorByID)

	// Inscripciones
	router.POST("/activities/:id/inscripciones", utils.CORS, controllers.CrearInscripcion)

	router.Run()
}
