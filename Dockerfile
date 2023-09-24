# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21.1-alpine3.18 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/simple-pass

RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-pass

WORKDIR /app

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /simple-auth /simple-auth

EXPOSE 8080

ENTRYPOINT ["/simple-pass"]