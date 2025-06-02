// Archivo: services/activities_services.go
package services

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"fmt"
)

// GetActivityByID busca en la base de datos una actividad por su ID.
// Devuelve domain.Activity con todos sus horarios, o error si no existe.
func GetActivityByID(id int) (domain.Activity, error) {
	// 1. Cargar el registro DAO de la actividad
	var daoAct dao.Actividad
	tx := clients.DB.First(&daoAct, id)
	if tx.Error != nil {
		return domain.Activity{}, fmt.Errorf("actividad con ID %d no encontrada: %w", id, tx.Error)
	}

	// 2. Cargar todos los horarios asociados a esta actividad
	var daoHorarios []dao.Horario
	tx = clients.DB.Where("actividad_id = ?", id).Find(&daoHorarios)
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

	// 4. Mapear dao.Actividad → domain.Activity
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
