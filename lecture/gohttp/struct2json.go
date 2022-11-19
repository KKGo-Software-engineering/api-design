package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{
		ID: 1, Name: "AnuchitO", Age: 18,
	}
	b, err := json.Marshal(u)
	fmt.Printf("byte : %s  \n", b)
	fmt.Println(err)
}
