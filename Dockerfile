# build stage
FROM golang:1.24.4 AS builder
WORKDIR /app
COPY . .
RUN go build cmd/main.go -o app

# final image
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/app .
COPY .env .
EXPOSE 8080
CMD ["./app"]