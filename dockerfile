# Builder
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency dan download
COPY go.mod go.sum ./
RUN go mod download

# Copy source code dan build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o backend-bengkel ./cmd/api

#  Runner
FROM alpine:3.22

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/backend-bengkel .

# Buka port 8600 sesuai Gin 
EXPOSE 8600

CMD ["./backend-bengkel"]