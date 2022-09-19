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

	stmt, err := db.Prepare("UPDATE users SET name=$2 WHERE id=$1;")

	if err != nil {
		log.Fatal("can't prepare statment update", err)
	}

	if _, err := stmt.Exec(1, "Aorjoa"); err != nil {
		log.Fatal("error execute update ", err)
	}

	fmt.Println("update success")

}
