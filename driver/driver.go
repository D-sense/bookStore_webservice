package driver

import (
	"database/sql"
	"github.com/d-sense/library_web_service/helper"
	"github.com/lib/pq"
	"os"
)


var db *sql.DB

func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	helper.LogFatal(err)

	db, err = sql.Open("postgres", pgURL)
	helper.LogFatal(err)

	err = db.Ping()
	helper.LogFatal(err)

	return db
}