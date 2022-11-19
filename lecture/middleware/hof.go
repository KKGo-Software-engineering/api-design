package main

import "fmt"

type Decorator func(s string) error

func Use(next Decorator) Decorator {
	return func(c string) error {
		fmt.Println("do something before")
		r := c + " should be green."
		return next(r)
	}
}

func home(s string) error {
	fmt.Println("home", s)
	return nil
}

func main() {
	wrapped := Use(home)
	w := wrapped("world")
	fmt.Println("end result:", w)
}
