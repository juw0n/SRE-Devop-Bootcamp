# Stage 1: Build environment (builder stage)
FROM golang:1.21.8-alpine3.18 AS builder-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/student-go-api && echo "Built application binary: docker-go-rest-api"

# Stage 2: Runtime environment
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder-stage /app/student-go-api .
COPY --from=builder-stage /app/app.env .

# Expose port 8000
EXPOSE 8000

# Run
CMD ["/app/student-go-api"]