package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/juw0n/SRE-Devop-Bootcamp/config"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	configVar, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("connot load config:", err)
	}
	conn, err := sql.Open(configVar.DBDriver, configVar.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
