# Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

# Copy dan download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY . .
COPY .env .env

# Build aplikasi utama (misal: main.go berada di ./cmd/api/main.go)
RUN go build -o app .

# Jalankan aplikasi
CMD ["./app", "start"]