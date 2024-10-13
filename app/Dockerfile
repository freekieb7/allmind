# syntax=docker/dockerfile:1 
FROM golang:1.23-alpine AS build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

RUN CGO_ENABLED=0 GOOS=linux go build -o /allmind
RUN CGO_ENABLED=0 GOOS=linux go build -C healthcheck -o /healthcheck

HEALTHCHECK --interval=5s --timeout=3s --start-period=2s --retries=3 CMD [ "/healthcheck" ]

# Run in development modes
FROM build as development

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go mod download

CMD ["air", "-c", ".air.toml"]

# Run the tests in the container
FROM build AS test
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS production

WORKDIR /

COPY --from=build /allmind /allmind
COPY --from=build /healthcheck /healthcheck

EXPOSE 3000

ENTRYPOINT ["/allmind"]