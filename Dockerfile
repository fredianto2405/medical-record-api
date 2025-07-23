# build stage
FROM golang:1.24.4 AS builder
WORKDIR /app

# copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build app
RUN go build -o app ./cmd/main.go

# run stage
FROM debian:bullseye-slim
WORKDIR /app

# copy the compiled binary from builder
COPY --from=builder /app/app .

# copy .env
COPY .env .

# expose port
EXPOSE 8080

# run app
CMD ["./app"]