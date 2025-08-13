# Vue 3 + Vite

This template should help get you started developing with Vue 3 in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

Learn more about IDE Support for Vue in the [Vue Docs Scaling up Guide](https://vuejs.org/guide/scaling-up/tooling.html#ide-support).

# test-TB

This is a fullstack application with a Vue.js frontend (Vite) and a Golang backend.

## Features

- User registration and login UI (Vue.js)
- Backend API for registration, login, password hashing, and JWT validation (Golang)
- Secure password storage and JWT-based authentication

## Getting Started

### Frontend

```bash
npm run dev
```

### Backend

```bash
cd backend
go run main.go
```

### Testing Backend

```bash
cd backend
go test -v
```

### Database

```bash
docker run --name testtb-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=testtb -p 5432:5432 -d postgres:latest
```

---

## Project Structure

- `/backend` - Golang backend API
- `/` - Vue.js frontend (Vite)

---

## To Do

- Implement registration and login UI
- Implement backend endpoints for registration, login, and JWT validation
- Connect frontend to backend


