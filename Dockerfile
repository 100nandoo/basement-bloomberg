# --- Stage 1: Build Frontend (Astro) ---
FROM node:lts-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build

# --- Stage 2: Build Backend (Go) ---
FROM golang:1.24-alpine AS backend-builder
WORKDIR /app/backend
# Install git for fetch dependencies
RUN apk add --no-cache git
# Copy source code and mod files
# We place go.mod at the root of /app/backend so it covers both cmd/ and go/
COPY backend/go/go.mod backend/go/go.sum ./
COPY backend/cmd/ ./cmd/
COPY backend/go/ ./go/

# Update dependencies (since we added new imports in main.go) and build
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# --- Stage 3: Final Image ---
FROM alpine:latest
WORKDIR /app

# Copy Backend Binary
COPY --from=backend-builder /app/backend/main .

# Copy Frontend Build Output
# Using 'dist' as the target directory expected by the Go app
COPY --from=frontend-builder /app/frontend/dist ./dist

EXPOSE 8080

CMD ["./main"]
