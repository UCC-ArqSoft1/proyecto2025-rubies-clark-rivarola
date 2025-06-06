// Archivo: services/activities_services.go
package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"fmt"
)

// GetActivityByID busca en la base de datos una actividad por su ID.
// Devuelve domain.Activity con todos sus horarios ordenados, o error si no existe.
func GetActivityByID(id int) (domain.Activity, error) {
	// 1. Cargar el registro DAO de la actividad
	var daoAct dao.Actividad
	tx := clients.DB.First(&daoAct, id)
	if tx.Error != nil {
		return domain.Activity{}, fmt.Errorf("actividad con ID %d no encontrada: %w", id, tx.Error)
	}

	// 2. Cargar todos los horarios asociados a esta actividad, ordenados correctamente
	var daoHorarios []dao.Horario
	tx = clients.DB.
		Where("actividad_id = ?", id).
		// Ordenamos por día de la semana en el orden Lun→Dom, y luego por hora_inicio
		Order("FIELD(dia_semana, 'Lun','Mar','Mie','Jue','Vie','Sab','Dom'), hora_inicio ASC").
		Find(&daoHorarios)
	if tx.Error != nil {
		return domain.Activity{}, fmt.Errorf("error cargando horarios para actividad %d: %w", id, tx.Error)
	}

	// 3. Mapear cada dao.Horario → domain.Horario
	var schedule []domain.Horario
	for _, h := range daoHorarios {
		schedule = append(schedule, domain.Horario{
			ID:          h.ID,
			ActividadID: h.ActividadID,
			Day:         h.DiaSemana,
			StartTime:   h.HoraInicio.Format("15:04"),
			DurationMin: h.DuracionMin,
		})
	}

	// 4. Mapear dao.Actividad → domain.Activity, incluyendo el arreglo de horarios ordenados
	domainAct := domain.Activity{
		ID:                  daoAct.ID,
		Name:                daoAct.Titulo,
		Description:         daoAct.Descripcion,
		Capacity:            daoAct.CupoMax,
		ImageURL:            daoAct.ImagenURL,
		CategoryName:        daoAct.NombreCategoria,
		InstructorName:      daoAct.NombreInstructor,
		InstructorEmail:     daoAct.EmailInstructor,
		InstructorSpecialty: daoAct.EspecialidadInstructor,
		Active:              daoAct.Activo,
		Schedule:            schedule,
	}

	return domainAct, nil
}

// ListarActividades devuelve un slice de ActivityListItem.
// Si `query` no es vacío, filtra por título o categoría.
func ListarActividades(query string) ([]domain.ActivityListItem, error) {
	var daoActs []dao.Actividad
	db := clients.DB

	// Aplicar filtro si se proporciona búsqueda
	if query != "" {
		like := "%" + query + "%"
		db = db.Where("titulo LIKE ? OR nombre_categoria LIKE ?", like, like)
	}

	// Obtener todas las actividades
	if tx := db.Find(&daoActs); tx.Error != nil {
		return nil, fmt.Errorf("error listando actividades: %w", tx.Error)
	}

	var resultados []domain.ActivityListItem
	for _, a := range daoActs {
		// Obtener el primer horario disponible para mostrar en el listado
		var h dao.Horario
		clients.DB.
			Where("actividad_id = ?", a.ID).
			Order("FIELD(dia_semana, 'Lun','Mar','Mie','Jue','Vie','Sab','Dom'), hora_inicio ASC").
			First(&h)

		var day, startTime string
		if h.ID != 0 {
			day = h.DiaSemana
			startTime = h.HoraInicio.Format("15:04")
		}

		item := domain.ActivityListItem{
			ID:             a.ID,
			Title:          a.Titulo,
			InstructorName: a.NombreInstructor,
			Day:            day,
			StartTime:      startTime,
		}
		resultados = append(resultados, item)
	}

	return resultados, nil
}
