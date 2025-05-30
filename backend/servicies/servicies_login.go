// Archivo: services/login_service.go
package services

import (
	"backend/clients"
	"backend/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) (uint, string, error) {
	var user models.Usuario

	err := clients.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return 0, "", fmt.Errorf("usuario no encontrado: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return 0, "", errors.New("contraseña incorrecta")
	}

	token, err := generateJWT(user)
	if err != nil {
		return 0, "", fmt.Errorf("error generando token: %w", err)
	}

	return user.ID, token, nil
}

func generateJWT(user models.Usuario) (string, error) {
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
