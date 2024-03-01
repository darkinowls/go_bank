migrateup:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose up

migratedown:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migratedown sqlc test