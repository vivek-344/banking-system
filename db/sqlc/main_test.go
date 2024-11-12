package db_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/vivek-344/banking-system/db/sqlc"
)

const (
	dbSource = "postgresql://root:vivek@localhost:5432/banking_system?sslmode=disable"
)

var testQueries *db.Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
