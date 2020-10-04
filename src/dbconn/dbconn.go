package dbconn

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
)

func Connect(dbUrl string) error {
	var err error
	Db, err = sql.Open("mysql", dbUrl)
	return err
}
