# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23-alpine AS build-stage

RUN apk --no-cache add \
    sqlite \
    gcc \
    g++ 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=1 GOOS=linux go build -o /allmind cmd/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /allmind /allmind

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]

# development stage
FROM build-stage AS build-development-stage

# RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", "third_party/.air.toml"]



# syntax=docker/dockerfile:1 
FROM golang:1.23-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go install github.com/a-h/templ/cmd/templ@latest && \
    templ generate

RUN CGO_ENABLED=1 GOOS=linux go build -C cmd/app -o tmp/allmind
RUN CGO_ENABLED=1 GOOS=linux go build -C cmd/healthcheck -o tmp/healthcheck

# Run in development modes
FROM build as development

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go mod download

HEALTHCHECK --interval=5s --timeout=3s --start-period=2s --retries=3 CMD [ "/healthcheck" ]

CMD ["air", "-c", "config/.air.toml"]

# Run the tests in the container
FROM build AS test
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS production

WORKDIR /

COPY --from=build /allmind /allmind
COPY --from=build /app/web/static /web/static
COPY --from=build /healthcheck /healthcheck

HEALTHCHECK --interval=5s --timeout=3s --start-period=2s --retries=3 CMD [ "/healthcheck" ]

EXPOSE 3000

ENTRYPOINT ["/allmind"]