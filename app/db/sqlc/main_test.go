package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://myuser:mypassword@localhost:5431/mydb?sslmode=disable"
)

var testQueries *Queries
var dbCon *sql.DB

func TestMain(m *testing.M) {
	var err error
	dbCon, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(dbCon)
	code := m.Run()
	os.Exit(code)
}
