package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ":", i)
	}
}

func main() {
	go slow("gopher")

	slow("nong")

	time.Sleep(10 * time.Second)
	fmt.Println("all task done.")
}
