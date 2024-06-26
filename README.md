# Reto Técnico Go y Node.js

Este proyecto contiene dos APIs implementadas en Go y Node.js, que trabajan en conjunto para recibir una matriz, rotarla y calcular estadísticas sobre la matriz rotada.

## Tecnologías Utilizadas

- Go (Golang)
- Node.js
- Docker
- Vue 3 (Frontend)
- JWT (Autenticación)

## Estructura del Proyecto

- `GOAPI`: Contiene la API de Go que rota la matriz.
- `nodeapi`: Contiene la API de Node.js que calcula estadísticas sobre la matriz rotada.
- `frontend`: Contiene la aplicación Vue 3 que interactúa con las APIs.

## Requisitos

- Docker
- Docker Compose
- Node.js
- Go (Golang)

## Instalación y Ejecución

### Paso 1: Clonar el Repositorio

git clone https://github.com/tu-usuario/tu-repositorio.git
cd tu-repositorio

### Paso 2: Construir y Ejecutar con Docker Compose

docker-compose up --build

### Paso 3: Probar las APIs con Postman

Obtener un token JWT:

URL: POST http://localhost:3000/login
Body (JSON):

{
    "username": "your-username"   EN ACTUALIZACION
}
Copia el token de la respuesta.
Rotar una matriz:

URL: POST http://localhost:8080/rotate
Headers: Authorization: Bearer your-jwt-token
Body (JSON):  LISTO

{
    "matrix": [
        [1, 2],
        [3, 4]
    ]
}
Calcular estadísticas:

URL: POST http://localhost:3000/stats
Headers: Authorization: Bearer your-jwt-token
Body (JSON):  LISTO

{
    "rotatedMatrix": [
        [3, 1],
        [4, 2]
    ]
}
Paso 4: Ejecutar el Frontend

cd frontend/matriz
npm install
npm run serve
Accede a http://localhost:8081 en tu navegador para ver la interfaz frontend.

Pruebas Unitarias
Para ejecutar las pruebas unitarias en Node.js, usa el siguiente comando:

npm test
