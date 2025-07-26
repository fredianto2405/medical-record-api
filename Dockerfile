# build stage (alphine)
FROM golang:1.24.4-alpine AS builder

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

# run stage (pakai base image yg cocok dgn GLIBC 2.34)
FROM debian:bookworm-slim
RUN apt update && apt install -y ca-certificates

WORKDIR /app
COPY --from=builder /app/app .
COPY .env .

EXPOSE 9000

CMD ["./app"]