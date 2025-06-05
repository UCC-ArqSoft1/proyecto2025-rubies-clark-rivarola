// Archivo: database.go
package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"backend/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		&dao.Usuario{},
		&dao.Actividad{},
		&dao.Inscripcion{},
	)
	if err != nil {
		log.Fatal("Error al migrar las tablas: ", err)
	}

	seedDB()
}

func seedDB() {

	// Crear usuario admin por defecto
	admin := dao.Usuario{
		Nombre:       "Administrador",
		Email:        "admin@gimnasio.com",
		PasswordHash: "admin123", // hashear antes en producción
		Rol:          "administrador",
		CreatedAt:    time.Now(),
	}
	DB.FirstOrCreate(&admin, dao.Usuario{Email: admin.Email})
}
