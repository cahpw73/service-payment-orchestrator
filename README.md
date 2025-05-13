# service-payment-orchestrator

Este proyecto es un microservicio escrito en Go que se encarga de **gestionar el proceso de pago de servicios** públicos o privados mediante la orquestación de varias operaciones necesarias para completar una transacción de pago.

## 🧠 Funcionalidades

El microservicio implementa diversas funcionalidades orientadas a facilitar la experiencia de pago de servicios y consultas asociadas:

- 📌 **Afiliaciones**: Registrar, actualizar o consultar las afiliaciones de usuarios a diferentes servicios.
- 🗂️ **Categorías y Subcategorías**: Consultar servicios organizados jerárquicamente por categoría y subcategoría.
- 🏙️ **Ciudades**: Gestión y consulta de ciudades disponibles para los servicios afiliados.
- 🧾 **Obtener deudas**: Consultar deudas activas asociadas a un servicio o cuenta afiliada.
- 💸 **Pagar deudas**: Realizar el pago de una o más deudas pendientes.
- 🔍 **Criterios de búsqueda**: Permite buscar servicios o deudas usando distintos filtros personalizados.
- 📑 **Servicios por subcategoría**: Obtener servicios específicos dentro de una subcategoría seleccionada.

## ⚙️ Configuración de entorno

El servicio utiliza variables de entorno configuradas en un archivo `.env` para facilitar la parametrización por ambiente. A continuación se listan y explican:

| Variable | Descripción |
|---------|-------------|
| `MIDDLEWARE_URL` | URL base del middleware o backend externo que se utiliza para validaciones, logins u otras integraciones |
| `MIDDLEWARE_SECRET` | Secreto usado para autenticación con el middleware (actualmente vacío) |
| `PORT` | Puerto en el que se expone el microservicio |
| `TOPAZ_CHANNEL` | Canal de comunicación para integraciones con Topaz (o proveedor externo) |
| `APPLICATION_ID` | ID único de la aplicación dentro del ecosistema |
| `DEVICE_ID` | Identificador lógico del dispositivo que hace la petición |
| `DEVICE_IP` | IP del dispositivo (usado en logging o auditoría) |
| `REDIS_HOST` | Dirección del servidor Redis para almacenamiento temporal |
| `REDIS_TTL` | Tiempo de vida (en minutos) para los datos almacenados en Redis |
| `DATABASE_CONEXION` | Cadena de conexión a base de datos MySQL |
| `ORACLE_TSNAME_DIR` | Ruta al directorio donde se encuentra el archivo `tnsnames.ora` del cliente Oracle |
| `ORACLE_CLIENT_BIN` | Ruta al cliente binario de Oracle Instant Client |

## 🐳 Docker

Este proyecto incluye un `Dockerfile` que permite contenerizar el microservicio para su despliegue en ambientes controlados (Dev, QA, Prod).

### 🔧 Comandos útiles

```bash
# Build del contenedor
docker build -t service-payment-orchestrator .

# Ejecutar el contenedor
docker run -d -p 8080:8080 --env-file .env --name orchestrator service-payment-orchestrator
```

## 📚 Tecnologías utilizadas

- **Go** como lenguaje principal
- **MySQL** para almacenamiento relacional
- **Redis** para almacenamiento en caché
- **Docker** para contenerización
- **Oracle Instant Client** (opcional) para integraciones con Oracle DB si aplica
- **HTTP y RESTful APIs** para integración con sistemas externos

## 📌 Notas adicionales

Este servicio se diseñó para actuar como **orquestador** dentro del ecosistema de pagos, lo cual significa que **coordina múltiples operaciones atómicas** (como obtener deuda, validar afiliación y realizar pago) para completar un flujo de negocio complejo de forma centralizada.

---

## 👨‍💻 Autor

**Christian Alba Herrera**  
Desarrollador de software especializado en backend, microservicios y arquitectura empresarial.  

📍 Bolivia

© 2025 – Todos los derechos reservados.