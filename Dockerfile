# build stage
FROM golang:1.24.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

# run stage (pakai base image yg cocok dgn GLIBC 2.34)
FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/app .
COPY .env .

EXPOSE 8080

CMD ["./app"]