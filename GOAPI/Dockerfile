# Usar la imagen base oficial de Go
FROM golang:1.22-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el código fuente de Go al directorio de trabajo del contenedor
COPY . .

# Descargar las dependencias necesarias especificadas en go.mod y go.sum
RUN go mod download

# Compilar la aplicación para producir un ejecutable binario
RUN go build -o main .

# Comando para ejecutar el binario compilado
CMD ["./main"]
