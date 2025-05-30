// Archivo: database.go
package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/models"
)

var DB *gorm.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando con la base de datos: ", err)
	}

	migrate()
}

func migrate() {
	err := DB.AutoMigrate(
		&models.Usuario{},
		&models.Categoria{},
		&models.Instructor{},
		&models.Actividad{},
		&models.Inscripcion{},
	)
	if err != nil {
		log.Fatal("Error al migrar las tablas: ", err)
	}

	seedDB()
}

func seedDB() {
	// Crear categorías de ejemplo
	categorias := []models.Categoria{
		{Nombre: "Yoga", Descripcion: "Actividades de relajación y flexibilidad"},
		{Nombre: "Funcional", Descripcion: "Entrenamiento físico integral"},
		{Nombre: "Spinning", Descripcion: "Entrenamiento aeróbico en bicicleta"},
	}
	for _, cat := range categorias {
		DB.FirstOrCreate(&cat, models.Categoria{Nombre: cat.Nombre})
	}

	// Crear instructores de ejemplo
	instructores := []models.Instructor{
		{Nombre: "Carlos Pérez", Email: "carlos@example.com", Especialidad: "Yoga"},
		{Nombre: "Lucía Gómez", Email: "lucia@example.com", Especialidad: "Spinning"},
	}
	for _, inst := range instructores {
		DB.FirstOrCreate(&inst, models.Instructor{Email: inst.Email})
	}

	// Crear usuario admin por defecto
	admin := models.Usuario{
		Nombre:       "Administrador",
		Email:        "admin@gimnasio.com",
		PasswordHash: "admin123", // hashear antes en producción
		Rol:          "administrador",
		CreatedAt:    time.Now(),
	}
	DB.FirstOrCreate(&admin, models.Usuario{Email: admin.Email})
}
