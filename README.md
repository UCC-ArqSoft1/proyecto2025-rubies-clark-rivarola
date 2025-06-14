# Projecto de Arquitectura de Softwere I
Autores: Clark Santiago, Rubies Catalina y Rivarola Valentina 


## 📋 Requisitos

- **MySQL** (≥ 8.0)
- **Go** (≥ 1.18)
- **Node.js** (≥ 16) y **npm** o **yarn**

-> Una vez clonado el proyecto, primero se debe: 
## 🗄️ 1. Preparar la base de datos

1. Crear la base de datos 
   - Abre MySQL Workbench o tu CLI MySQL.
   - Ejecuta dentro de MySQL:
     ```sql
     CREATE DATABASE gimnasio_actividades;
     ```
2. Importar tablas y datos de ejemplo 
   - Desde MySQL Workbench:  
     - Archivo → Run SQL Script… → selecciona `db.sql` 
3. Verificar 
   - Comprueba que las tablas `usuarios`, `actividades`, `horarios` e `inscripciones` existen y tienen filas de ejemplo.

Luego en la terminal entrar al proycto y a la carpeta de backend y pegar: 
# En VS Code Terminal (PowerShell), pegá estas líneas:
$Env:DB_HOST     = "127.0.0.1"
$Env:DB_PORT     = "3306"
$Env:DB_USER     = "root"
$Env:DB_PASSWORD = "TU_CONTRASEÑA" //(Contraseña de la base de datos) 
$Env:DB_NAME     = "backend"

# Ahora ejecutá tu servidor:
go run .\main.go

Si todo va bien verás en consola:
[GIN-DEBUG] Listening and serving HTTP on :8080

RUTAS DISPONIBLES:
->POST /users/login
->GET /activities
->GET /activities/:id
->POST /inscripciones

## 🌐 Configurar y levantar el frontend: 

npm install //muy importante para que se instale de node_modules en el frontend(nosotros lo hacemos desde la terminal de Node.js command prompt)
Desde la carpeta frontend/ en la terminal: 
npm run de node (tambien lo hacemos desde la terminal de Node.js command prompt)
Abre tu navegador en http://localhost:5173 (o el puerto que indique Vite).

Para probar el login se pueden usar los siguintes usuarios: 
Usuario: admin y Contraseña: Password123
Usuario: emiliano y Contraseña: Password123
Usuario: maria y Contraseña: Password123
(Si la contraseña o el usuarios son incorrectos devolvera "credenciales inválidas")

En la pantalla de Home se puede probarque si el usuario ya esta inscripto en la actividad 
el programa devolvera "ya estás inscrito a este horario", por lo contrario si el usuario no esta inscriprto 
en la actividad y se preciona el boton de inscribirse devolverav "¡Inscripción exitosa!". 

4. Probar la aplicación
Login
POST http://localhost:8080/users/login
Body JSON:
{ "username": "usuario1", "password": "tuPassword" }

Listar actividades
GET http://localhost:8080/activities

Detalle de actividad
GET http://localhost:8080/activities/1

Inscribirse
POST http://localhost:8080/inscripciones
Body JSON:
{ "usuario_id": 1, "horario_id": 10 }

En el frontend, iniciá sesión en la pantalla de login y luego verás el listado de actividades. Podrás:
Filtrar por título, categoría, día u horario.
Inscribirte pulsando “Inscribirse” junto a cada actividad.