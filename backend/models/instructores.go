// Archivo: models/instructor.go
package models

type Instructor struct {
	ID           uint16  `gorm:"primaryKey;column:id;autoIncrement"`
	Nombre       string  `gorm:"column:nombre;type:varchar(100);not null"`
	Email        string  `gorm:"column:email;type:varchar(150);unique"`
	Especialidad *string `gorm:"column:especialidad;type:varchar(100)"`

	Actividades []Actividad `gorm:"foreignKey:InstructorID"`
}

func (Instructor) TableName() string {
	return "instructores"
}
