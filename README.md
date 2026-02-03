# Basement Bloomberg

This is a simple stock tracking application. It features a Go backend that provides a REST API for stock data and a modern frontend built with Astro.

The backend is a simple Go server using the `chi` router, and the frontend is built with Astro. The two services are intended to be run concurrently during development.

## Project Structure

```text
/
├── backend/
│   └── cmd/
│       └── api/
│           └── main.go
├── frontend/
│   ├── public/
│   │   ├── favicon.ico
│   │   └── favicon.svg
│   ├── src/
│   │   └── pages/
│   │       └── index.astro
│   ├── .gitignore
│   ├── astro.config.mjs
│   ├── package.json
│   ├── tsconfig.json
│   └── ...
└── readme.md
```