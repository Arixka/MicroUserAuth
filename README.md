
# microUserAuth

## Descripción
**microUserAuth** es un microservicio en Go diseñado para la autenticación y gestión de usuarios. Este servicio utiliza el framework Gin para crear una API RESTful y PostgreSQL como base de datos, empaquetado en Docker para facilitar su despliegue en entornos de desarrollo y producción. Este proyecto se centra en la seguridad y eficiencia en el manejo de sesiones y contraseñas.

## Características
- **Registro de Usuarios**: Permite a los usuarios registrarse con un nombre de usuario y una contraseña segura.
- **Autenticación**: Maneja el inicio de sesión de usuarios, generando tokens de autenticación (JWT) para acceso seguro.
- **Almacenamiento Seguro de Contraseñas**: Las contraseñas se almacenan de forma segura utilizando técnicas de hashing.
- **Integración con Docker**: Configuración de Docker y Docker Compose para simplificar el despliegue y administración de entornos.
- **Documentación de API**: Documentación de los endpoints para facilitar la integración con otros servicios.

## Tecnologías Utilizadas
- **Go**: Lenguaje principal para el desarrollo del servicio.
- **Gin Framework**: Framework ligero y rápido para crear aplicaciones web en Go.
- **PostgreSQL**: Base de datos relacional utilizada para almacenar la información de los usuarios.
- **Docker y Docker Compose**: Para contenedorización del servicio y administración de dependencias.

## Configuración y Uso

### Requisitos Previos
- **Go** (versión 1.x o superior)
- **Docker** y **Docker Compose**

### Instalación y Ejecución
1. **Clonar el repositorio**:

   ```bash
   git clone https://github.com/tuusuario/microUserAuth.git
   ```

2. **Navegar al directorio del proyecto**:

   ```bash
   cd microUserAuth
   ```

3. **Iniciar los servicios con Docker Compose**:

   ```bash
   docker-compose up
   ```

4. **Variables de Entorno**: Asegúrate de configurar las variables de entorno necesarias en un archivo `.env`. Aquí tienes un ejemplo de las variables más comunes:

   ```
   DB_HOST=db
   DB_USER=usuario
   DB_PASSWORD=contraseña
   DB_NAME=nombre_base_de_datos
   JWT_SECRET=clave_secreta_para_jwt
   ```

### Ejecución usando Makefile

Este proyecto incluye un `Makefile` con varios comandos para simplificar el proceso de desarrollo y despliegue:

- **Compilar el binario**:

  ```bash
  make build
  ```

- **Ejecutar el servicio en local** (sin Docker):

  ```bash
  make run-local
  ```

- **Iniciar el servicio y la base de datos con Docker Compose**:

  ```bash
  make run-docker
  ```

- **Detener el servicio y eliminar contenedores**:

  ```bash
  make down-docker
  ```

Estos comandos permiten un manejo más sencillo del proyecto durante el desarrollo y despliegue. Asegúrate de configurar las variables de entorno necesarias antes de ejecutar los comandos.

### Documentación de la API
La API expone varios endpoints para el registro y autenticación de usuarios. Aquí tienes un resumen de los principales:

- **POST /users**: Registra un nuevo usuario.
  - **Cuerpo**: `{ "username": "string", "password": "string" }`
  - **Respuesta**: JSON con los datos del usuario registrado o un mensaje de error.

- **POST /login**: Autentica a un usuario y devuelve un token JWT.
  - **Cuerpo**: `{ "username": "string", "password": "string" }`
  - **Respuesta**: `{ "token": "jwt_token" }` en caso de éxito, o un mensaje de error.

## Licencia
Este proyecto está licenciado bajo la MIT License - ver el archivo [LICENSE](LICENSE) para más detalles.

## Contacto

Para consultas o más información, puedes contactarme en:
[![Gmail](https://img.shields.io/badge/-Gmail-c14438?style=flat&logo=Gmail&logoColor=white)](mailto:marisiver25@gmail.com)
[![Linkedin](https://img.shields.io/badge/-LinkedIn-blue?style=flat&logo=Linkedin&logoColor=white)](https://www.linkedin.com/in/maria-siverio/)
