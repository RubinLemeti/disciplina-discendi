# Stage 1: Build the Go application
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .

# Stage 2: Create a lightweight runtime container
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["/app/main"]
