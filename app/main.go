package main

import (
	"app/api"
	db "app/db/sqlc"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable"
	serverAddress = ":8081"
)

//var testQueries *Queries
//var dbCon *sql.DB

func main() {
	dbCon, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(dbCon)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
