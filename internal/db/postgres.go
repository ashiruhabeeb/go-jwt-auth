package db

import(
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(driver, cred string) (*Database, error){
	db, err := sql.Open(driver, cred)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) CloseDB() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
