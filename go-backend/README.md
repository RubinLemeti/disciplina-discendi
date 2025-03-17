# Go backend
This backend will handle auth and user CRUD. The idea is to have a REST API for User CRUD and possibly a gRPC for auth.

## Database migrations

Database migrations can be created with the command below
```sh
migrate create -ext sql -dir ./internal/db/migrations create_schema 
```

The migrations can be run with the command below
```sh
migrate -source file:///app/internal/db/migrations -database "postgres://root:root@postgres_db:5432/go_backend?sslmode=disable" up
```