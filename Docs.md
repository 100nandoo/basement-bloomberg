# Development Docs

## Libraries
* FlyonUI
* Highcharts

## Sturcture
```text
my-project/
├── backend/                # Go (Chi) source code
│   ├── cmd/
│   │   └── api/            # Entry point: main.go starts the server
│   ├── internal/           # Private code (not importable by other projects)
│   │   ├── api/            # Chi routes and middleware
│   │   ├── service/        # Business logic (Usecases)
│   │   ├── repository/     # Data storage logic (DB queries)
│   │   └── model/          # Domain entities/structs
│   ├── config/             # Config loading (env/yaml)
│   └── go.mod
├── frontend/               # Astro project (added later)
│   ├── src/
│   ├── public/
│   └── package.json
├── docker/                 # Dockerfiles and orchestration
│   ├── backend.Dockerfile
│   └── frontend.Dockerfile
├── .env                    # Environment variables
└── docker-compose.yml      # Orchestrates Go, Astro, and Cloudflared
```