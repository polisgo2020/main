FROM golang:1.14 as builder

WORKDIR /usr/src

COPY go.mod .
COPY go.sum .
RUN go mod download

# При изменении проекта все строки выше не будут выполняться, если не изменились зависимости в файлах `go.sum`, `go.mod`.
COPY . .
# Флаги нужны, чтоб приложение запускалось в образе alpine:latest
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o blog

# alpine -- минималистичный linux дистрибутив
FROM alpine:latest
RUN apk update && apk add --no-cache tzdata ca-certificates

WORKDIR /usr/app
COPY --from=builder /usr/src/blog .
COPY --from=builder /usr/src/view ./view
CMD ["./blog"]
