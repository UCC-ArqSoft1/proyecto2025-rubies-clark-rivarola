package domain

// User se usa para binding de login y, opcionalmente, para representar datos de usuario.
type User struct {
	ID       uint   `json:"id,omitempty"`       // Se incluye al devolver información del usuario (p. ej., en perfiles).
	Username string `json:"username"`           // Nombre de usuario para login.
	Password string `json:"password,omitempty"` // Contraseña en texto plano solo para binding; no devolver al cliente.
	Email    string `json:"email,omitempty"`    // Correo del usuario, opcionalmente útil en otros contextos.
	Role     string `json:"role,omitempty"`     // Rol del usuario ("socio" o "administrador") si se necesita exponer.
}
