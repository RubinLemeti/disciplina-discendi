#sh
migrate -source file:///app/internal/db/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@postgres_db:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}" up