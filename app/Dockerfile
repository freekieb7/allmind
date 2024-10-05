# syntax=docker/dockerfile:1 
FROM golang:1.23-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /allmind

# Run the tests in the container
FROM build AS test
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM scratch AS prod

WORKDIR /

COPY --from=build /allmind /allmind

EXPOSE 3000

ENTRYPOINT ["/allmind"]




# RUN go install github.com/air-verse/air@latest
# RUN go install github.com/go-delve/delve/cmd/dlv@latest
# RUN go install github.com/a-h/templ/cmd/templ@latest

# CMD ["air", "-c", ".air.toml"]

# HEALTHCHECK --interval=5m --timeout=3s \
#   CMD curl -f http://localhost:$HTTP_PORT/debug/vars || exit 1