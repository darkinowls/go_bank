migrate_up:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose up

migrate_down:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrate_up migrate_down sqlc test