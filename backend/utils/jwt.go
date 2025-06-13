package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtDuration = time.Hour * 24
	jwtSecret   = "jwtSecret" // Valor por defecto; en producción usar variable de entorno
)

// GenerateJWT genera un token JWT con RegisteredClaims.
// Ahora acepta userID como uint para coincidir con dao.Usuario.ID.
func GenerateJWT(userID uint) (string, error) {
	// Setear expiración
	expirationTime := time.Now().Add(jwtDuration)

	// Construir los claims registrados
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "backend",
		Subject:   "auth",
		ID:        fmt.Sprintf("%d", userID),
	}

	// Crear el token con HS256 y los claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token usando el secreto
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}

// ValidateJWT verifica un token JWT y devuelve el userID extraído de los RegisteredClaims.
func ValidateJWT(tokenString string) (uint, error) {
	// Parsear y validar firma
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el método de firma sea HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method unexpected")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("token inválido o expirado")
	}

	// Extraer los RegisteredClaims
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return 0, errors.New("error al extraer claims")
	}

	// El campo ID en RegisteredClaims se guardó como string del userID
	if claims.ID == "" {
		return 0, errors.New("user_id no encontrado en token")
	}
	// Convertir ID string a uint
	var userID uint
	_, err = fmt.Sscanf(claims.ID, "%d", &userID)
	if err != nil {
		return 0, errors.New("error convirtiendo user_id")
	}

	return userID, nil
}

// AuthMiddleware es un middleware de Gin que:
// - Extrae el header "Authorization: Bearer <token>".
// - Valida el JWT y, si es válido, añade "userID" al contexto.
// - Si falla, retorna 401 Unauthorized y aborta la ejecución.
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token faltante o mal formado"})
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		userID, err := ValidateJWT(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		// Guardar userID en el contexto
		ctx.Set("userID", userID)
		ctx.Next()
	}
}
