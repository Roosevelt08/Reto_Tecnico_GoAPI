# RETO_TECNICO_GOAPI

Este proyecto consiste en una API desarrollada en Go y Node.js para la rotación de matrices y el cálculo de estadísticas, así como un frontend en Vue.js para interactuar con estas APIs. Además, el proyecto está configurado para ejecutarse utilizando Docker y Docker Compose.

## Estructura del Proyecto

El proyecto tiene la siguiente estructura:


RETO_TECNICO_GOAPI
├── frontend
│   └── matriz
│       ├── node_modules
│       ├── public
│       │   ├── favicon.ico
│       │   └── index.html
│       ├── src
│       │   ├── assets
│       │   │   └── logo.png
│       │   ├── components
│       │   │   ├── HelloWorld.vue
│       │   │   ├── Login.vue
│       │   │   ├── MatrixDisplay.vue
│       │   │   ├── MatrixInput.vue
│       │   │   └── StatisticsDisplay.vue
│       │   ├── router
│       │   │   └── index.js
│       │   ├── views
│       │   │   ├── AboutView.vue
│       │   │   └── HomeView.vue
│       │   ├── App.vue
│       │   └── main.js
│       ├── .gitignore
│       ├── babel.config.js
│       ├── jsconfig.json
│       ├── package.json
│       ├── package-lock.json
│       ├── README.md
│       └── vue.config.js
├── GOAPI
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── GoAPI.exe
│   ├── main.go
│   └── node_modules
└── nodeapi
    ├── node_modules
    ├── app.js
    ├── app.test.js
    ├── Dockerfile
    ├── matrixUtils.js
    ├── package-lock.json
    ├── package.json
    └── docker-compose.yml


## Configuración y Ejecución

### Prerrequisitos

- Docker
- Docker Compose
- Node.js
- npm

### Instrucciones

1. Clona el repositorio:

   sh
   git clone https://github.com/Roosevelt08/Reto_Tecnico_GoAPI.git
   cd RETO_TECNICO_GOAPI
   

2. Navega al directorio `frontend/matriz` y ejecuta:

   npm install
   

3. Navega al directorio raíz del proyecto y ejecuta Docker Compose para construir y levantar los servicios:

   docker-compose up --build
   

### Descripción de las APIs

#### API en Go (GOAPI)

- **Endpoint:** `POST /rotate`
- **Descripción:** Rota una matriz dada.
- **Ejemplo de cuerpo de solicitud:**

  json
  {
    "matrix": [
      [1, 2],
      [3, 4]
    ]
  }
  

- **Ejemplo de respuesta:**

  json
  {
    "rotatedMatrix": [
      [3, 1],
      [4, 2]
    ],
    "maxValue": 4,
    "minValue": 1,
    "average": 2.5,
    "totalSum": 10,
    "isDiagonal": false
  }
  

#### API en Node.js (nodeapi)

- **Endpoint:** `POST /rotate`
- **Descripción:** Rota una matriz dada.
- **Ejemplo de cuerpo de solicitud:**

  json
  {
    "matrix": [
      [1, 2],
      [3, 4]
    ]
  }
  

- **Ejemplo de respuesta:**

  json
  {
    "rotatedMatrix": [
      [3, 1],
      [4, 2]
    ],
    "maxValue": 4,
    "minValue": 1,
    "average": 2.5,
    "totalSum": 10,
    "isDiagonal": false
  }
  

### Frontend en Vue.js

- **Descripción:** Interfaz gráfica para interactuar con las APIs.
- **Instrucciones:** 
  1. Navega al directorio `frontend/matriz`.
  2. Ejecuta el servidor de desarrollo con `npm run serve`.
  3. Abre el navegador en `http://localhost:8081`.

## Docker

El proyecto está configurado para ejecutarse con Docker Compose. Asegúrate de tener Docker y Docker Compose instalados y ejecuta:

docker-compose up --build


Esto levantará los servicios definidos en el archivo `docker-compose.yml`.

## Ejecución de Pruebas

### Go
Para ejecutar las pruebas en Go, usa el siguiente comando en el terminal:

go test


### Node.js
Para ejecutar las pruebas en Node.js, usa el siguiente comando en el terminal dentro del directorio `nodeapi`:

npm test

