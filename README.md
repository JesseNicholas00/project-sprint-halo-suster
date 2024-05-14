Data Migration Script
migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path migrations up