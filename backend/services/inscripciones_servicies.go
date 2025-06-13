package services

import (
	"backend/clients"
	"backend/dao"
	"errors"
	"fmt"
)

// CrearInscripcion permite que un usuario se inscriba a un horario específico.
func CrearInscripcion(usuarioID, horarioID uint) error {
	// 1. Verificar que el horario exista y esté asociado a una actividad activa
	var horario dao.Horario
	if err := clients.DB.Preload("Actividad").First(&horario, horarioID).Error; err != nil {
		return fmt.Errorf("horario no encontrado")
	}
	if !horario.Actividad.Activo {
		return errors.New("la actividad a la que pertenece este horario está inactiva")
	}

	// 2. Contar cuántos usuarios ya están inscritos a este horario
	var conteo int64
	if err := clients.DB.Model(&dao.Inscripcion{}).
		Where("horario_id = ?", horarioID).
		Count(&conteo).Error; err != nil {
		return fmt.Errorf("error contando inscripciones: %w", err)
	}

	if uint16(conteo) >= horario.Actividad.CupoMax {
		return errors.New("no hay cupos disponibles para este horario")
	}

	// 3. Verificar si ya está inscrito
	var existente dao.Inscripcion
	if err := clients.DB.
		Where("usuario_id = ? AND horario_id = ?", usuarioID, horarioID).
		First(&existente).Error; err == nil {
		return errors.New("ya estás inscrito a este horario")
	}

	// 4. Crear la inscripción
	insc := dao.Inscripcion{
		UsuarioID: usuarioID,
		HorarioID: horarioID,
	}
	if err := clients.DB.Create(&insc).Error; err != nil {
		return fmt.Errorf("error creando la inscripción: %w", err)
	}

	return nil
}

// ListarInscripcionesPorUsuario devuelve los horarios a los que un usuario está inscrito.
func ListarInscripcionesPorUsuario(usuarioID uint) ([]dao.Horario, error) {
	var horarios []dao.Horario

	// Join entre inscripciones y horarios
	err := clients.DB.
		Joins("JOIN inscripciones ON inscripciones.horario_id = horarios.id").
		Where("inscripciones.usuario_id = ?", usuarioID).
		Preload("Actividad").
		Find(&horarios).Error

	if err != nil {
		return nil, fmt.Errorf("error consultando inscripciones del usuario: %w", err)
	}

	return horarios, nil
}
