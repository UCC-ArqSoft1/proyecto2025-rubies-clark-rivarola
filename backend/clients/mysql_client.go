// Archivo: clients/mysql_clients.go
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
	dsn := fmt.Sprintf(dsnFormat, "root", "root", "127.0.0.1", 3306, "backend")

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}

	// Migrar todas las tablas necesarias: Usuario, Actividad, Horario, Inscripcion
	tables := []interface{}{
		&dao.Usuario{},
		&dao.Actividad{},
		&dao.Horario{},
		&dao.Inscripcion{},
	}

	for _, table := range tables {
		if err := DB.AutoMigrate(table); err != nil {
			panic(fmt.Errorf("error migrating table: %w", err))
		}
	}

}

// GetUserByUsername busca un usuario por su nombre (campo 'nombre').
func GetUserByUsername(username string) (dao.Usuario, error) {
	var userDAO dao.Usuario
	txn := DB.First(&userDAO, "nombre = ?", username)
	if txn.Error != nil {
		return dao.Usuario{}, fmt.Errorf("error getting user: %w", txn.Error)
	}
	return userDAO, nil
}
