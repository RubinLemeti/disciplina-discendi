# Stage 1: Build the Go application
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
# RUN apt-get update \
# && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN apt-get update \
&& go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# creating the DB
CMD ["sh ./create-db.sh"]
RUN CGO_ENABLED=0 GOOS=linux go mod tidy && go build -o main ./cmd

CMD ["/app/main"]

# Stage 2: Create a lightweight runtime container
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder /app/main .
# CMD ["/app/main"]
