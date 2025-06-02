// Archivo: dao/inscripcion.go
package dao

import "time"

// Inscripcion representa la inscripción de un usuario a un horario específico.
type Inscripcion struct {
	ID               uint      `gorm:"primaryKey;column:id;autoIncrement"`
	UsuarioID        uint      `gorm:"column:usuario_id;not null"`
	HorarioID        uint      `gorm:"column:horario_id;not null"`
	FechaInscripcion time.Time `gorm:"column:fecha_inscripcion;autoCreateTime"`

	// Relaciones
	Usuario Usuario `gorm:"foreignKey:UsuarioID;references:ID"`
	Horario Horario `gorm:"foreignKey:HorarioID;references:ID"`
}

func (Inscripcion) TableName() string {
	return "inscripciones"
}
