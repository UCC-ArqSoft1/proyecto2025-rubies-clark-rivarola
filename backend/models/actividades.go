// Archivo: models/actividad.go
package models

import (
	"time"
)

type Actividad struct {
	ID          uint      `gorm:"primaryKey;column:id;autoIncrement"`
	Titulo      string    `gorm:"column:titulo;not null"`
	Descripcion string    `gorm:"column:descripcion;type:text;not null"`
	DiaSemana   string    `gorm:"column:dia_semana;type:enum('Lun','Mar','Mie','Jue','Vie','Sab','Dom');not null"`
	HoraInicio  time.Time `gorm:"column:hora_inicio;type:time;not null"`
	DuracionMin uint16    `gorm:"column:duracion_min;not null"`
	CupoMax     uint16    `gorm:"column:cupo_max;not null"`
	ImagenURL   *string   `gorm:"column:imagen_url;type:varchar(255)"`

	CategoriaID  uint16 `gorm:"column:categoria_id"`
	InstructorID uint16 `gorm:"column:instructor_id"`

	Categoria  Categoria  `gorm:"foreignKey:CategoriaID;references:ID"`
	Instructor Instructor `gorm:"foreignKey:InstructorID;references:ID"`

	Inscripciones []Inscripcion `gorm:"foreignKey:ActividadID"`
}

func (Actividad) TableName() string {
	return "actividades"
}
