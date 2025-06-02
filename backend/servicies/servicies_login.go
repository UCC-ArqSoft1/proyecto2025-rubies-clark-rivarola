// Archivo: services/login_service.go
package services

import (
	"backend/clients"
	"backend/dao"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Login valida credenciales y retorna el ID del usuario y un JWT si son correctas.
func Login(username string, password string) (uint, string, error) {
	var user dao.Usuario

	// Buscamos por el campo "nombre" en la tabla usuarios (que mapea a Username en DAO)
	err := clients.DB.Where("nombre = ?", username).First(&user).Error
	if err != nil {
		return 0, "", fmt.Errorf("usuario no encontrado")
	}

	// Comparamos la contraseña hasheada almacenada con la ingresada
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return 0, "", errors.New("credenciales inválidas")
	}

	// Generamos JWT
	token, err := generateJWT(user)
	if err != nil {
		return 0, "", fmt.Errorf("error generando token: %w", err)
	}

	return user.ID, token, nil
}

// generateJWT crea un token JWT con claims básicos (user_id, username, rol, exp)
func generateJWT(user dao.Usuario) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secreto123" // valor por defecto para pruebas
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"rol":      user.Rol,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
