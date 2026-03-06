# 1. IMAGEN BASE: Linux ligero y estable para entorno industrial
FROM python:3.11-slim-bookworm

# 2. METADATOS: Sello de Propiedad de Edmundo Ramírez
LABEL maintainer="Edmundo Ramirez"
LABEL project="CORE-COLD-V2-INDUSTRIAL"

# 3. HERRAMIENTAS DE HIERRO: Instalamos compiladores de C++ y dependencias de red
RUN apt-get update && apt-get install -y \
    g++ \
    gcc \
    libc6-dev \
    make \
    && rm -rf /var/lib/apt/lists/*

# 4. ENTORNO DE TRABAJO
WORKDIR /usr/src/core_cold

# 5. INYECCIÓN DE CÓDIGO: Pasamos todo el arsenal al contenedor
COPY . .

# 6. FORJA DEL GATEKEEPER: Compilamos el centinela C++
RUN g++ src/gatekeeper/gatekeeper.cpp -o gatekeeper_bin

# 7. EJECUCIÓN: Iniciamos la Capa 1 por defecto
# CORE-COLD arranca protegiendo el perímetro
CMD ["./gatekeeper_bin"]
