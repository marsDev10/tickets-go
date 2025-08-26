# Dockerfile para desarrollo
FROM golang:1.24.3-alpine AS development

# Instalar dependencias del sistema
RUN apk add --no-cache git curl

# Instalar Air para hot reload en desarrollo
RUN go install github.com/air-verse/air@latest

# Configurar directorio de trabajo
WORKDIR /app

# Copiar archivos de dependencias primero (para aprovechar cache de Docker)
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Crear directorio tmp para Air
RUN mkdir -p tmp

# Exponer puerto
EXPOSE 8080

# Comando por defecto para desarrollo
CMD ["air", "-c", ".air.toml"]

# ============================================
# Stage de producción
# ============================================
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Copiar el resto del código
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage para producción
FROM alpine:latest AS production

WORKDIR /app

# Instalar certificados SSL para HTTPS requests
RUN apk --no-cache add ca-certificates

# Copiar solo el binario compilado
COPY --from=builder /app/main .

# Exponer puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]