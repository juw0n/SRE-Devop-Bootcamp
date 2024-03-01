# syntax=docker/dockerfile:1

FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-api

EXPOSE 8080

# Run
CMD ["/docker-go-api"]