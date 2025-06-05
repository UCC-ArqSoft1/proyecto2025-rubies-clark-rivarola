// Archivo: domain/horario.go
package domain

// Horario representa un turno específico de una actividad.
type Horario struct {
	ID          uint   `json:"id"`
	ActividadID uint   `json:"actividad_id"`
	Day         string `json:"day"`        // "Lun", "Mar", "Mie", "Jue", "Vie", "Sab", "Dom"
	StartTime   string `json:"start_time"` // Formato "HH:MM"
	DurationMin uint16 `json:"duration"`   // Duración en minutos
}
