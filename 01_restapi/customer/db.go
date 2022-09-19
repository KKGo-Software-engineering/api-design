package customer

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/proullon/ramsql/driver"
)

var db *sql.DB

func init() {
	conn, err := sql.Open(os.Getenv("DATABASE_DRIVER"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	db = conn
	createCustomerTable()
}

func createCustomerTable() {
	createCustomerTable := `
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`

	if _, err := db.Exec(createCustomerTable); err != nil {
		log.Fatal("can't create table ", err)
	}
}
