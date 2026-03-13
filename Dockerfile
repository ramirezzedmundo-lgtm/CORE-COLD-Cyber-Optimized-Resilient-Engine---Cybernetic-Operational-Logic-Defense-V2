# ETAPA 1: El Taller (Compilación)
FROM golang:1.21-alpine AS builder

# Configuración de entorno para ejecución en "Hierro" (Estructura estática)
ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod ./
# Si tienes dependencias, descomenta la siguiente línea:
# RUN go mod download

# Copiar el código fuente que acabamos de refinar
COPY main.go .

# Compilar un binario estático y optimizado
RUN go build -ldflags="-s -w" -o bridge main.go

# ETAPA 2: La Cápsula (Ejecución Ligera)
FROM scratch

WORKDIR /

# Copiar solo el binario desde el "Taller"
COPY --from=builder /app/bridge /bridge

# Exponer el puerto del Bridge para comunicación interna (L3)
EXPOSE 18181

# Ejecutar el centinela
ENTRYPOINT ["/bridge"]
