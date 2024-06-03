FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./quake-parser-log ./cmd

FROM alpine
COPY --from=build /app/quake-parser-log /usr/local/bin/app
COPY --from=build /app/static/html /static/html

ENTRYPOINT ["app", "-f", "/static/html"]

EXPOSE 8080