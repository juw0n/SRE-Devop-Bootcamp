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
RUN go build -o student-go-api && echo "Built application binary: student-go-rest-api"
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz

# Stage 2: Runtime environment
FROM alpine:latest

RUN apk update && apk add bash

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder-stage /app/student-go-api .
COPY --from=builder-stage /app/app.env .
COPY --from=builder-stage /app/start.sh .
COPY --from=builder-stage /app/migrate ./migrate
COPY --from=builder-stage /app/database/migration ./migration

# Change permissions of start.sh to make it executable
RUN chmod +x start.sh

# Expose port 8000
EXPOSE 8000

# Run
ENTRYPOINT [ "/app/start.sh" ]
CMD [ "/app/student-go-api" ]