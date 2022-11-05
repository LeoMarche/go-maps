package sqloader

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenGPKG(filename string) (*sql.DB, error) {
	return sql.Open("sqlite3", filename)
}
