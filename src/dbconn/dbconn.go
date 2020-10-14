package dbconn

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
)

func Connect(dbUrl string) error {
	Db, _ = sql.Open("mysql", dbUrl)

	err := Db.Ping()
	return err
}
