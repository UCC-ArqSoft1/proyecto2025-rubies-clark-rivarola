-- Descripción de db.sql: Crea la base de datos, las tablas y carga datos.

-- Crear la base de datos si no existe

CREATE DATABASE IF NOT EXISTS backend
  DEFAULT CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

USE backend;

-- Tabla: usuarios
--------------------------------------------
CREATE TABLE IF NOT EXISTS usuarios (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    password_hash VARCHAR(256) NOT NULL,
    rol ENUM('socio','administrador') NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

-- Tabla: actividades
--------------------------------------------
CREATE TABLE IF NOT EXISTS actividades (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(100) NOT NULL,
    descripcion TEXT NOT NULL,
    cupo_max SMALLINT NOT NULL,
    imagen_url VARCHAR(255),
    nombre_categoria VARCHAR(50) NOT NULL,
    nombre_instructor VARCHAR(100) NOT NULL,
    email_instructor VARCHAR(150),
    especialidad_instructor VARCHAR(100),
    activo TINYINT(1) NOT NULL DEFAULT 1,
    INDEX idx_actividades_categoria (nombre_categoria)
) ENGINE=InnoDB;

-- Tabla: horarios (cada fila es un turno concreto de una actividad)
--------------------------------------------
CREATE TABLE IF NOT EXISTS horarios (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    actividad_id BIGINT NOT NULL,
    dia_semana ENUM('Lun','Mar','Mie','Jue','Vie','Sab','Dom') NOT NULL,
    hora_inicio TIME NOT NULL,
    duracion_min SMALLINT NOT NULL,
    CONSTRAINT fk_horarios_actividades FOREIGN KEY (actividad_id)
        REFERENCES actividades(id)
        ON DELETE CASCADE,
    INDEX idx_horarios_actividad (actividad_id),
    INDEX idx_horarios_dia_semana (dia_semana)
) ENGINE=InnoDB;

-- Tabla: inscripciones
--------------------------------------------
CREATE TABLE IF NOT EXISTS inscripciones (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    usuario_id BIGINT NOT NULL,
    horario_id BIGINT NOT NULL,
    fecha_inscripcion DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(usuario_id, horario_id),
    CONSTRAINT fk_inscripciones_usuarios FOREIGN KEY (usuario_id)
        REFERENCES usuarios(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_inscripciones_horarios FOREIGN KEY (horario_id)
        REFERENCES horarios(id)
        ON DELETE CASCADE,
    INDEX idx_inscripciones_usuario (usuario_id),
    INDEX idx_inscripciones_horario (horario_id)
) ENGINE=InnoDB;

-- ----------------------------------------------------------------------------
--                  INSERTAMOS DATOS EN LA BASE DE DATOS                     --
-- ----------------------------------------------------------------------------

-- Insertamos usuarios d
--------------------------------------------
-- El hash bcrypt corresponde a la contraseña "password123"
INSERT INTO usuarios (nombre, email, password_hash, rol) VALUES
  ('emiliano', 'emiliano@mail.com', '$2b$12$Kzai13XwycEQFCgCySDAjeaPtRTmHpFaCJzJGISytzlskjFkbzhvO', 'socio'),
  ('maria',    'maria@mail.com',    '$2b$12$Kzai13XwycEQFCgCySDAjeaPtRTmHpFaCJzJGISytzlskjFkbzhvO', 'socio'),
  ('admin',    'admin@gimnasio.com','$2b$12$Kzai13XwycEQFCgCySDAjeaPtRTmHpFaCJzJGISytzlskjFkbzhvO', 'administrador');

-- Insertar actividades 
--------------------------------------------
INSERT INTO actividades 
  (titulo, descripcion, cupo_max, imagen_url, nombre_categoria, nombre_instructor, email_instructor, especialidad_instructor, activo)
VALUES
  ('Zumba Avanzado',
   'Clase intensa de Zumba combinando ritmos latinos y cardio.',
   20,
   'https://example.com/img/zumba.jpg',
   'Baile',
   'Florencia Ruiz',
   'florencia@instructores.com',
   'Zumba certificada',
   TRUE),

  ('Funcional Básico',
   'Entrenamiento funcional para todos los niveles, usa peso corporal.',
   15,
   NULL,
   'Fitness',
   'Carlos Pérez',
   'carlos@instructores.com',
   'Entrenador Personal',
   TRUE),

  ('Pilates Suave',
   'Rutina suave de Pilates para mejorar la flexibilidad y el core.',
   10,
   'https://example.com/img/pilates.jpg',
   'Pilates',
   'Ana Gómez',
   NULL,
   NULL,
   TRUE);

-- Insertamos horarios relacionados a cada actividad
--------------------------------------------
-- Para 'Zumba Avanzado' (actividad_id = 1)
INSERT INTO horarios (actividad_id, dia_semana, hora_inicio, duracion_min) VALUES
  (1, 'Lun', '18:00:00', 60),
  (1, 'Mie', '19:00:00', 60),
  (1, 'Vie', '17:00:00', 60);

-- Para 'Funcional Básico' (actividad_id = 2)
INSERT INTO horarios (actividad_id, dia_semana, hora_inicio, duracion_min) VALUES
  (2, 'Mar', '18:30:00', 45),
  (2, 'Jue', '19:30:00', 45);

-- Para 'Pilates Suave' (actividad_id = 3)
INSERT INTO horarios (actividad_id, dia_semana, hora_inicio, duracion_min) VALUES
  (3, 'Lun', '10:00:00', 50),
  (3, 'Mie', '10:00:00', 50);

-- Insertamos algunas inscripciones para empezar
-- 'emiliano' (usuario_id = 1) se inscribe en 'Zumba Avanzado' lunes 18:00 (horario_id = 1)
INSERT INTO inscripciones (usuario_id, horario_id) VALUES
  (1, 1),

-- 'emiliano' también en 'Funcional Básico' martes 18:30 (horario_id = 4)
  (1, 4),

-- 'maria' (usuario_id = 2) se inscribe en 'Pilates Suave' lunes 10:00 (horario_id = 7)
  (2, 7);


-- ------------------------------------------------------------------
--        Inserta usuarios adicionales, actividades y horarios     --
--          para enriquecer los datos de prueba del backend.       --
-- ------------------------------------------------------------------

-- Usuarios adicionales
--------------------------------------------
INSERT INTO usuarios (nombre, email, password_hash, rol) VALUES
  ('juanperez',  'juan@mail.com',  '$2b$12$Kzai13XwycEQFCgCySDAjeaPtRTmHpFaCJzJGISytzlskjFkbzhvO', 'socio'),
  ('luisa',      'luisa@mail.com', '$2b$12$Kzai13XwycEQFCgCySDAjeaPtRTmHpFaCJzJGISytzlskjFkbzhvO', 'socio'),
  ('pabloh',     'pablo@mail.com', '$2b$12$Kzai13XwycEQFCgCySDAjeaPtRTmHpFaCJzJGISytzlskjFkbzhvO', 'socio');

-- Actividades nuevas
--------------------------------------------
INSERT INTO actividades 
  (titulo, descripcion, cupo_max, imagen_url, nombre_categoria, nombre_instructor, email_instructor, especialidad_instructor, activo)
VALUES
  ('Zumba Intermedio',
   'Clase de Zumba para nivel intermedio, con coreografía desafiante.',
   25,
   'https://statics-cuidateplus.marca.com/cms/images/zumba.jpg',
   'Baile',
   'María López',
   'maria@instructores.com',
   'Ritmos Latinos',
   TRUE),

  ('Funcional Intensivo',
   'Entrenamiento funcional de alta intensidad para quemar calorías.',
   20,
   NULL,
   'Fitness',
   'Carlos Pérez',
   'carlos@instructores.com',
   'CrossFit Trainer',
   TRUE),

  ('Calistenia Básica',
   'Introducción a la calistenia: uso del peso corporal.',
   15,
   'https://example.com/img/calistenia.jpg',
   'Calistenia',
   'Ana Gómez',
   NULL,
   NULL,
   TRUE),

  ('Spinning Avanzado',
   'Clase de ciclismo indoor con intervalos de alta resistencia.',
   30,
   'https://example.com/img/spinning.jpg',
   'Ciclismo',
   'Roberto Díaz',
   'roberto@instructores.com',
   'Instructor Spinning',
   TRUE);

-- 3) Horarios para las nuevas actividades
--------------------------------------------
INSERT INTO horarios (actividad_id, dia_semana, hora_inicio, duracion_min) VALUES
  (4, 'Mar', '18:00:00', 60),
  (4, 'Jue', '19:00:00', 60);

-- Funcional Intensivo (actividad_id = 5)
INSERT INTO horarios (actividad_id, dia_semana, hora_inicio, duracion_min) VALUES
  (5, 'Lun', '20:00:00', 50),
  (5, 'Mie', '20:00:00', 50),
  (5, 'Vie', '20:00:00', 50);

-- Calistenia Básica (actividad_id = 6)
INSERT INTO horarios (actividad_id, dia_semana, hora_inicio, duracion_min) VALUES
  (6, 'Mar', '10:00:00', 45),
  (6, 'Jue', '10:00:00', 45);

-- Spinning Avanzado (actividad_id = 7)
INSERT INTO horarios (actividad_id, dia_semana, hora_inicio, duracion_min) VALUES
  (7, 'Mié', '07:00:00', 45),
  (7, 'Sab', '09:00:00', 50);

