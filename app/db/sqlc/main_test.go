package db

import (
	"app/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var dbCon *sql.DB

func TestMain(m *testing.M) {
	conf, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return
	}
	dbCon, err = sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(dbCon)
	code := m.Run()
	os.Exit(code)
}
