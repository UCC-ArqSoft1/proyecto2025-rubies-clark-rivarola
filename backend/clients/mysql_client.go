package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	dsnFormat := "%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local"
	dsn := fmt.Sprintf(dsnFormat, "root", "Santi020923", "127.0.0.1", 3306, "backend")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}

	// Ejecutar AutoMigrate para TODAS las entidades del proyecto
	err = DB.AutoMigrate(
		&dao.Usuario{},
		&dao.Actividad{},
		&dao.Horario{},
		&dao.Inscripcion{},
	)
	if err != nil {
		panic(fmt.Errorf("error en AutoMigrate: %w", err))
	}
}

// GetUserByUsername busca un usuario por su nombre de usuario (campo nombre)
func GetUserByUsername(username string) (dao.Usuario, error) {
	var user dao.Usuario
	txn := DB.Where("nombre = ?", username).First(&user)
	if txn.Error != nil {
		return dao.Usuario{}, fmt.Errorf("error al buscar usuario: %w", txn.Error)
	}
	return user, nil
}
