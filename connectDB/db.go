package connectdb

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=pguser password=pgpassword dbname=sorteador sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	return db, err
}
