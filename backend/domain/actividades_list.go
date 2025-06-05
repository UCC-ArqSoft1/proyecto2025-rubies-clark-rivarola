package domain

// ActivityListItem representa la información mínima de una actividad
// para mostrarse en el listado o búsqueda.
type ActivityListItem struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	InstructorName string `json:"instructor_name"`
	Day            string `json:"day"`        // Ej: "Lun", "Mar", …
	StartTime      string `json:"start_time"` // Formato "HH:MM"
}
