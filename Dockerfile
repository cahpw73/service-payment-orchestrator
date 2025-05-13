# Usando la imagen base de Go
FROM golang:1.23-bullseye AS builder

# Copiar archivos de la aplicación y construir el binario
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o service-payment-orchestrator

# Etapa final
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y libaio1 wget unzip
RUN apt-get install -y build-essential

# Variables de entorno para Oracle en la etapa final
ENV ORACLE_BASE=/usr/lib/instantclient
ENV TNS_ADMIN=/usr/lib/instantclient
ENV ORACLE_HOME=/usr/lib/instantclient
ENV TZ=America/La_Paz

# Estableciendo directorio de trabajo
WORKDIR /tmp

# Instalar tzdata (datos de zona horaria)
RUN apt-get install -y tzdata

# Instalar libaio, libnsl y compatibilidad con glibc (en lugar de libc6-compat)
RUN apt-get install -y libaio1 libnsl-dev

WORKDIR /usr/lib

# Descargar e instalar Oracle Instant Client
RUN wget https://download.oracle.com/otn_software/linux/instantclient/2360000/instantclient-basic-linux.x64-23.6.0.24.10.zip && \
    unzip instantclient-basic-linux.x64-23.6.0.24.10.zip && rm -f instantclient-basic-linux.x64-23.6.0.24.10.zip && \
    mv instantclient_23_6 instantclient && \
    cd instantclient && rm -f *jdbc* *occi* *mysql* *mql1* *ipc1* *jar uidrvci genezi adrci && \
    echo "/usr/lib/instantclient" > /etc/ld.so.conf.d/oracle-instantclient.conf && ldconfig

# Configurar la variable de entorno
ENV LD_LIBRARY_PATH=/usr/lib/instantclient

# Establecer directorio de trabajo
WORKDIR /usr/src/app

# Copiar bibliotecas de Oracle, el binario y el archivo .env
COPY --from=builder /usr/src/app/service-payment-orchestrator /usr/src/app/service-payment-orchestrator
COPY .env ./.env

# Configurar zona horaria
RUN apt-get update && apt-get install -y --no-install-recommends tzdata && \
    echo "America/La_Paz" > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata

# Exponer el puerto de la aplicación (ajustar según necesidad)
EXPOSE 8080

# Comando de inicio
CMD ["./service-payment-orchestrator"]


