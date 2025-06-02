// Archivo: services/users_services.go
package services

import (
	"errors"
	"fmt"

	"backend/clients"
	"backend/dao"
	"backend/utils"

	"golang.org/x/crypto/bcrypt"
)

// Login autentica a un usuario comparando su contraseña y, de ser correcta,
// genera y devuelve un JWT junto al ID del usuario.
func Login(username string, password string) (uint, string, error) {
	// 1. Obtener al usuario por su nombre (campo `nombre` en la tabla `usuarios`)
	var user dao.Usuario
	if tx := clients.DB.Where("nombre = ?", username).First(&user); tx.Error != nil {
		return 0, "", fmt.Errorf("usuario no encontrado")
	}

	// 2. Comparar la contraseña ingresada con el hash almacenado en la base
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return 0, "", errors.New("credenciales inválidas")
	}

	// 3. Generar JWT utilizando el utilitario
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return 0, "", fmt.Errorf("error generando token: %w", err)
	}

	return user.ID, token, nil
}
