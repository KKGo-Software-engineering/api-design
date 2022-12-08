package main

import "fmt"

type Database struct {
	conn string
}

func (d *Database) Close() {
	d.conn = "close"
	fmt.Println("database connection is closed")
}

func (d *Database) Exec() {
	if d.conn == "close" {
		fmt.Println("database connection is closed, can't exec")
		return
	}

	fmt.Println("exec success")
}

func nong() string {
	fmt.Println("defer nong")
	fmt.Println("nong")

	return "hello gopher"
}

func main() {
	nong()
	
	// conn := Database{}

	// conn.Exec()
	// // error

	// conn.Close()
}
