FROM golang:1.14 as builder

WORKDIR /usr/src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o blog

FROM alpine:latest
RUN apk update && apk add --no-cache tzdata ca-certificates

WORKDIR /usr/app
COPY --from=builder /usr/src/blog .
COPY --from=builder /usr/src/view ./view
CMD ["./blog"]
