package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2)  RETURNING id", "Aorjoa", 19)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't insert data", err)
	}
	fmt.Println("insert todo success id : ", id)
}
