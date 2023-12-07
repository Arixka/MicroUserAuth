# Etapa de construcción
FROM golang:1.21.4-alpine AS build_base

# Instala las dependencias necesarias
RUN apk add --no-cache git gcc g++

# Configura el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código fuente
COPY . .

# Compila la aplicación Go
RUN go build -o ./out/microUserAuth ./cmd/api/main.go

# Etapa final usando una imagen más pequeña
FROM alpine:latest

# Instala las dependencias necesarias en el entorno de ejecución
RUN apk --no-cache add ca-certificates

# Copia el ejecutable compilado desde la etapa de construcción
COPY --from=build_base /app/out/microUserAuth /app/microUserAuth

# Expone el puerto 8080
EXPOSE 8080

# Comando para ejecutar tu aplicación
CMD ["/app/microUserAuth"]
