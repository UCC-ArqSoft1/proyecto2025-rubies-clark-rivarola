-- Base de datos: sistema de actividades deportivas para gimnasio

-- Tabla: usuarios
CREATE TABLE IF NOT EXISTS usuarios (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    password_hash VARCHAR(256) NOT NULL,
    rol ENUM('socio', 'administrador') NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

-- Tabla: categorias
CREATE TABLE IF NOT EXISTS categorias (
    id SMALLINT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL UNIQUE,
    descripcion VARCHAR(200)
) ENGINE=InnoDB;

-- Tabla: instructores
CREATE TABLE IF NOT EXISTS instructores (
    id SMALLINT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE,
    especialidad VARCHAR(100)
) ENGINE=InnoDB;

-- Tabla: actividades
CREATE TABLE IF NOT EXISTS actividades (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(100) NOT NULL,
    descripcion TEXT NOT NULL,
    dia_semana ENUM('Lun', 'Mar', 'Mie', 'Jue', 'Vie', 'Sab', 'Dom') NOT NULL,
    hora_inicio TIME NOT NULL,
    duracion_min SMALLINT NOT NULL,
    cupo_max SMALLINT NOT NULL,
    imagen_url VARCHAR(255),
    categoria_id SMALLINT,
    instructor_id SMALLINT,
    FOREIGN KEY (categoria_id) REFERENCES categorias(id),
    FOREIGN KEY (instructor_id) REFERENCES instructores(id)
) ENGINE=InnoDB;

-- Tabla: inscripciones
CREATE TABLE IF NOT EXISTS inscripciones (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    usuario_id BIGINT NOT NULL,
    actividad_id BIGINT NOT NULL,
    fecha_inscripcion DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id),
    FOREIGN KEY (actividad_id) REFERENCES actividades(id),
    UNIQUE(usuario_id, actividad_id) -- evita inscripciones duplicadas
) ENGINE=InnoDB;
