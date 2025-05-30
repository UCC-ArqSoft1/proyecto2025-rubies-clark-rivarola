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

	tables := []interface{}{
		&models.Usuario{},
		&models.Categoria{},
		&models.Instructor{},
		&models.Actividad{},
		&models.Inscripcion{},
	}

	for _, table := range tables {
		if err := DB.AutoMigrate(table); err != nil {
			panic(fmt.Errorf("error migrating table: %w", err))
		}
	}

}

func GetUserByUsername(username string) (models.Usuario, error) {
	var user models.Usuario
	txn := DB.First(&user, "username = ?", username)
	if txn.Error != nil {
		return models.Usuario{}, fmt.Errorf("error getting user: %w", txn.Error)
	}
	return user, nil
}
