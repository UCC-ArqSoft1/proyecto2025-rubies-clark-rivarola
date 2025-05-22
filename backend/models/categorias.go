// Archivo: models/categoria.go
package models

type Categoria struct {
	ID          uint16  `gorm:"primaryKey;column:id;autoIncrement"`
	Nombre      string  `gorm:"column:nombre;type:varchar(50);not null;unique"`
	Descripcion *string `gorm:"column:descripcion;type:varchar(200)"`

	Actividades []Actividad `gorm:"foreignKey:CategoriaID"`
}

func (Categoria) TableName() string {
	return "categorias"
}
