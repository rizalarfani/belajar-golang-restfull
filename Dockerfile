# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.24.6 AS builder
WORKDIR /app

# Leverage Go modules caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary for Linux without CGO
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/restful-api ./

# Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app

# Copy compiled binary and any needed assets
COPY --from=builder /app/bin/restful-api ./restful-api

# Expose application port
EXPOSE 3000

# Run the service
ENTRYPOINT ["/app/restful-api"]
