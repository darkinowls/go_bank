package main

import (
	"app/api"
	db "app/db/sqlc"
	"app/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return
	}
	dbCon, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(dbCon)
	server := api.NewServer(store)

	err = server.Start(conf.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
