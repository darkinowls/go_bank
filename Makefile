migrate_up:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose up

migrate_down:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: migrate_up, migrate_down, sqlc