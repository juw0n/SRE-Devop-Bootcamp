# Stage 1: Build environment (builder stage)
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy all project files
COPY . .

# Build the application binary with a descriptive output
RUN go build -o go-rest-api && echo "Built application binary: go-rest-api"

# Stage 2: Final image (runtime stage)
FROM postgres:latest

WORKDIR /app

# Copy the application binary from the builder stage
COPY --from=builder ./app/go-rest-api /app

# Expose the port for your API
EXPOSE 8000

# Inject environment variables from Docker runtime
ENV POSTGRES_HOST=${POSTGRES_HOST}
ENV POSTGRES_PORT=${POSTGRES_PORT}
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}

# Set the application command using environment variables
CMD ["sh", \
    "-c", \    
    "go-rest-api", \
  "-db-host", "$POSTGRES_HOST", \
  "-db-port", "$POSTGRES_PORT", \
  "-db-user", "$POSTGRES_USER", \
  "-db-password", "$POSTGRES_PASSWORD" \
]
