migratecreate:
	migrate create -ext sql -dir ./app/db/migrations -seq $(ARGS)

migrateup:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose up

migrateup1:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
	cd app && mockgen -package mockdb -destination db/mock/store.go app/db/sqlc Store

test:
	cd app && go test -v -cover ./...

dcu:
	docker-compose up -d

server:
	cd app && go run main.go

.PHONY: migrateup migratedown sqlc test dcu server migratecreate