package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jainam1259/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
