**Helpdesk Tickets (Docker)**

- Backend en `Go (Gin)` con recarga en caliente usando `air`.
- Frontend `React + Vite + TypeScript`.
- Base de datos `PostgreSQL` y Adminer para administración.

**Requisitos**

- `Docker` y `Docker Compose` instalados.
- No necesitas Go ni Node instalados para desarrollo; los contenedores los proveen.

**Inicio Rápido**

- Copia el repo y, desde la raíz, ejecuta:
  - `docker compose up -d --build`
- Accesos por defecto:
  - Frontend: `http://localhost:5173`
  - Backend (API): `http://localhost:8080`
  - Adminer (DB): `http://localhost:8081`
    - Sistema: `PostgreSQL`
    - Servidor: `postgres`
    - Usuario: `postgres`
    - Contraseña: `password123`
    - Base de datos: `helpdesk`

**Servicios (docker-compose)**

- `postgres`: `postgres:15-alpine`, puerto `5432`, datos en volumen `postgres_data`.
- `backend`: Go (Gin) con hot reload vía `air`, expone `8080`.
- `frontend`: Vite dev server, expone `5173`.
- `adminer`: cliente web para DB, expone `8081`.

Revisa `docker-compose.yml` para más detalles y personalización.

**Variables de Entorno**

- Backend (`backend` en `docker-compose.yml`):
  - `DB_HOST=postgres`, `DB_PORT=5432`, `DB_USER=postgres`, `DB_PASSWORD=password123`, `DB_NAME=helpdesk`
  - `JWT_SECRET=your-super-secret-jwt-key-change-in-production`
  - `GIN_MODE=debug`, `PORT=8080`, `CGO_ENABLED=0`
- Frontend (`frontend`):
  - `VITE_API_URL=http://localhost:8080`
  - `NODE_ENV=development`

Puedes ajustar estos valores directamente en `docker-compose.yml` o moverlos a un `.env` y referenciarlos desde Compose.

**Desarrollo**

- Levantar entorno con recarga en caliente:
  - `docker compose up -d --build`
- Ver logs:
  - Backend: `docker compose logs -f backend`
  - Frontend: `docker compose logs -f frontend`
  - Base de datos: `docker compose logs -f postgres`
- El backend usa `air` con configuración en `backend/.air.toml`.
- El frontend monta el código con volumen y ejecuta `npm run dev` con `--host 0.0.0.0`.

**Producción (build de imágenes)**

- Backend (stage `production` de `backend/Dockerfile`):
  - `docker build -t helpdesk-backend:prod --target production backend`
- Frontend (stage `production` de `frontend/Dockerfile`):
  - `docker build -t helpdesk-frontend:prod --target production frontend`
  - Nota: el `Dockerfile` de frontend copia `frontend/nginx.conf`; crea ese archivo si vas a servir con Nginx en producción.
- Ejecución ejemplo (puertos locales):
  - Backend: `docker run -p 8080:8080 --env-file <tu-env> helpdesk-backend:prod`
  - Frontend: `docker run -p 80:80 helpdesk-frontend:prod`

Para una orquestación de producción, considera crear `docker-compose.prod.yml` con estos targets y variables seguras.

**Estructura**

- `docker-compose.yml`: definición de servicios de desarrollo.
- `backend/`: API en Go
  - `Dockerfile`, `main.go`, `routes/`, `controllers/`, `models/`, `db/connection.go`, `.air.toml`.
- `frontend/`: Vite React + TS
  - `Dockerfile`, `src/`, `vite.config.ts`, `package.json`.
- `database/init/`: scripts SQL opcionales que se ejecutan al iniciar Postgres (si existen).

**Comandos Útiles**

- Reconstruir y levantar: `docker compose up -d --build`
- Apagar: `docker compose down`
- Apagar y borrar volúmenes (incluye datos de DB): `docker compose down -v`
- Shell en un contenedor: `docker compose exec backend sh` (o `frontend`/`postgres`)

**Notas y Consejos**

- Los valores de DB y `JWT_SECRET` son solo para desarrollo. Cambia credenciales en producción.
- Si `frontend` no alcanza al backend, revisa `VITE_API_URL` y CORS en la API.
- Si hay conflicto de puertos, ajusta los mapeos en `docker-compose.yml`.

**Licencia**

- MIT ( Prueba )

