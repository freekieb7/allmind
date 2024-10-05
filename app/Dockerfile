# syntax=docker/dockerfile:1
FROM golang:1.23-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download



RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

CMD ["air", "-c", ".air.toml"]