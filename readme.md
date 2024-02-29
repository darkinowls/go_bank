## Migrations

```bash
migrate create -ext sql -dir ./app/db/migrations -seq init
```

```bash
migrate -path ./app/db/migrations -database "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable" -verbose up
```

```bash
sqlc generate
```

### deps

- sqlc
- golang-migrate