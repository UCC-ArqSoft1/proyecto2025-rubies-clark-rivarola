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

	// Buscamos por el campo "nombre" en la tabla usuarios
	if tx := clients.DB.Where("nombre = ?", username).First(&user); tx.Error != nil {
		return 0, "", fmt.Errorf("usuario no encontrado")
	}

	// Comparamos la contraseña ingresada con el hash almacenado
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return 0, "", errors.New("credenciales inválidas")
	}

	// Generamos JWT
	token, err := generateJWT(user.ID)
	if err != nil {
		return 0, "", fmt.Errorf("error generando token: %w", err)
	}

	return user.ID, token, nil
}

// generateJWT crea un token JWT con claims básicos (user_id, exp).
func generateJWT(userID uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secreto123" // valor por defecto para pruebas
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
