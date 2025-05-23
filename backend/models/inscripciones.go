// Archivo: models/inscripcion.go
package models

import "time"

type Inscripcion struct {
	ID               uint      `gorm:"primaryKey;column:id;autoIncrement"`
	UsuarioID        uint      `gorm:"column:usuario_id;not null"`
	ActividadID      uint      `gorm:"column:actividad_id;not null"`
	FechaInscripcion time.Time `gorm:"column:fecha_inscripcion;autoCreateTime"`

	Usuario   Usuario   `gorm:"foreignKey:UsuarioID;references:ID"`
	Actividad Actividad `gorm:"foreignKey:ActividadID;references:ID"`
}

func (Inscripcion) TableName() string {
	return "inscripciones"
}
