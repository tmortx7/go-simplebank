package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/tmortx7/go-simplebank/config"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config.ReadConfig(config.ReadConfigOption{})

	var err error

	testDB, err = sql.Open(config.C.Database.DBDriver, config.C.Database.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
