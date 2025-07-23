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
	// Carga el registro DAO de la actividad
	var daoAct dao.Actividad
	tx := clients.DB.First(&daoAct, id)
	if tx.Error != nil {
		return domain.Activity{}, fmt.Errorf("actividad con ID %d no encontrada: %w", id, tx.Error)
	}

	// Carga todos los horarios asociados a esta actividad, ordenados correctamente
	var daoHorarios []dao.Horario
	tx = clients.DB.
		Where("actividad_id = ?", id).
		// Ordenamos por día de la semana y luego por hora_inicio
		Order("FIELD(dia_semana, 'Lun','Mar','Mie','Jue','Vie','Sab','Dom'), hora_inicio ASC").
		Find(&daoHorarios)
	if tx.Error != nil {
		return domain.Activity{}, fmt.Errorf("error cargando horarios para actividad %d: %w", id, tx.Error)
	}

	// Mapea cada dao.Horario → domain.Horario
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

	// 4. Mapea dao.Actividad → domain.Activity, incluyendo el arreglo de horarios ordenados
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

func ListarActividades(query string) ([]domain.ActivityListItem, error) {
	var daoActs []dao.Actividad
	db := clients.DB

	if query != "" {
		like := "%" + query + "%"
		db = db.Where("titulo LIKE ? OR nombre_categoria LIKE ?", like, like)
	}

	if tx := db.Find(&daoActs); tx.Error != nil {
		return nil, fmt.Errorf("error listando actividades: %w", tx.Error)
	}

	var resultados []domain.ActivityListItem
	for _, a := range daoActs {
		// traigo el primer horario para mostrar día, hora y duración
		var h dao.Horario
		clients.DB.
			Where("actividad_id = ?", a.ID).
			Order("FIELD(dia_semana, 'Lun','Mar','Mie','Jue','Vie','Sab','Dom'), hora_inicio ASC").
			First(&h)

		// preparo valores por defecto
		day, start, duration := "", "", uint16(0)
		if h.ID != 0 {
			day = h.DiaSemana
			start = h.HoraInicio.Format("15:04")
			duration = h.DuracionMin // <-- ahora viene de h, no de a
		}

		resultados = append(resultados, domain.ActivityListItem{
			ID:             a.ID,
			Title:          a.Titulo,
			CategoryName:   a.NombreCategoria,
			InstructorName: a.NombreInstructor,
			Day:            day,
			StartTime:      start,
			Duration:       duration,
			Capacity:       a.CupoMax,
			ImageURL:       a.ImagenURL,
		})
	}

	return resultados, nil
}
