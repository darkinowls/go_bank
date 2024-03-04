migrateup:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose up

migratedown:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	cd app && mockgen -package mockdb -destination db/mock/store.go app/db/sqlc Store

test:
	cd app && go test -v -cover ./...

dcu:
	docker-compose up -d

server:
	go run app/main.go

.PHONY: migrateup migratedown sqlc test dcu server