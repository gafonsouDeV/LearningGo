to run migrations
migrate -path ./migrations -database "postgres://admin:admin123@localhost:5432/expense_tracker_db?sslmode=disable" down
migrate -path ./migrations -database "postgres://admin:admin123@localhost:5432/expense_tracker_db?sslmode=disable" up