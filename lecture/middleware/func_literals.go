package main

import "fmt"

func main() {
	r := func(a, b int) bool {
		return a < b
	}(2, 3)

	fmt.Println("result:", r)
}
