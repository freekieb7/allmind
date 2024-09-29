FROM golang:1.23-alpine

RUN go env -w CGO_ENABLED=1

# RUN apk update && apk add --no-cache build-base cmake git
RUN apk update && apk upgrade && apk add --no-cache build-base sqlite

COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod download

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

CMD ["air", "-c", ".air.toml"]